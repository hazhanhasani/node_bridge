package rest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/hazhanhasani/node_bridge/common"
	"github.com/hazhanhasani/node_bridge/controller"
	"github.com/hazhanhasani/node_bridge/tools"
)

type Node struct {
	controller.Controller
	client     *http.Client
	ctx        context.Context
	baseUrl    string
	cancelFunc context.CancelFunc
	mu         sync.Mutex
}

func New(address string, port int, serverCA []byte, apiKey uuid.UUID, logChanSize int, extra map[string]interface{}) (*Node, error) {
	certPool, err := tools.LoadClientPool(serverCA)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Node{
		Controller: controller.New(apiKey, logChanSize, extra),
		client: tools.CreateHTTPClient(certPool, address),
		ctx: ctx,
		baseUrl: "https://" + net.JoinHostPort(address, fmt.Sprintf("%d", port)),
		cancelFunc: cancel,
	}, nil
}

func (n *Node) Start(config string, backendType common.BackendType, users []*common.User, keepAlive uint64, excludeInbounds ...string) error {
	if n.Health() != controller.NotConnected {
		n.Stop()
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	data := &common.Backend{Type: backendType, Config: config, Users: users, KeepAlive: keepAlive, ExcludeInbounds: excludeInbounds}
	n.client.Timeout = 15 * time.Second
	var info common.BaseInfoResponse
	if err := n.createRequest(n.client, http.MethodPost, "start", data, &info); err != nil {
		return err
	}
	n.Connect(info.GetNodeVersion(), info.GetCoreVersion())
	n.client.Timeout = 10 * time.Second
	n.ctx, n.cancelFunc = context.WithCancel(context.Background())
	n.StartSync(n.ctx, n.SyncUsers)
	return nil
}

func (n *Node) Stop() {
	if n.Health() == controller.NotConnected {
		return
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	n.cancelFunc()
	n.Disconnect()
	_ = n.createRequest(n.client, http.MethodPut, "stop", &common.Empty{}, &common.Empty{})
}

func (n *Node) Info() (*common.BaseInfoResponse, error) {
	var info common.BaseInfoResponse
	if err := n.createRequest(n.client, http.MethodGet, "info", &common.Empty{}, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (n *Node) createRequest(client *http.Client, method, endpoint string, data proto.Message, response proto.Message) error {
	body, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, n.baseUrl+"/"+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("x-api-key", n.ApiKey())
	req.Header.Set("Content-Type", "application/x-protobuf")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	responseBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(responseBody))
	}
	if len(responseBody) == 0 {
		return nil
	}
	return proto.Unmarshal(responseBody, response)
}

func (n *Node) createStreamingRequest(client *http.Client, method, endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, n.baseUrl+"/"+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", n.ApiKey())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return resp.Body, nil
}
