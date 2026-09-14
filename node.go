package bluepanel_node_bridge

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/hazhanhasani/node_bridge/common"
	"github.com/hazhanhasani/node_bridge/controller"
	"github.com/hazhanhasani/node_bridge/rest"
	"github.com/hazhanhasani/node_bridge/rpc"
)

type BluePanelNode interface {
	Start(config string, backendType common.BackendType, users []*common.User, keepAlive uint64, excludeInbounds ...string) error
	Stop()
	NodeVersion() string
	CoreVersion() string
	SyncUsers(users []*common.User) error
	Info() (*common.BaseInfoResponse, error)
	GetSystemStats() (*common.SystemStatsResponse, error)
	GetBackendStats() (*common.BackendStatsResponse, error)
	GetStats(reset bool, name string, statType common.StatType) (*common.StatResponse, error)
	GetOutboundsLatency(name string) (*common.LatencyResponse, error)
	GetUserOnlineStat(string) (*common.OnlineStatResponse, error)
	GetUserOnlineIpList(string) (*common.StatsOnlineIpListResponse, error)
	ListRoutingRules() (*common.RoutingRulesResponse, error)
	GetBalancerInfo(tag string) (*common.BalancerInfoResponse, error)
	TestRoute(request *common.TestRouteRequest) (*common.RouteResult, error)
	AddRoutingRule(rule string, shouldReset bool) error
	RemoveRoutingRule(ruleTag string) error
	OverrideBalancerTarget(balancerTag, target string) error
	Health() controller.Health
	UpdateUsers([]*common.User)
	StreamLogs(context.Context) (<-chan controller.LogEntry, error)
	HardReset() <-chan struct{}
}

type NodeProtocol string

const (
	GRPC NodeProtocol = "GRPC"
	REST NodeProtocol = "REST"
)

type NodeOptions struct {
	address      string
	port         int
	serverCA     []byte
	apiKey       uuid.UUID
	extra        map[string]interface{}
	nodeProtocol NodeProtocol
	logChanSize  int
}

type NodeOption func(*NodeOptions) error

func WithPort(port int) NodeOption {
	return func(opts *NodeOptions) error {
		if port <= 0 {
			return errors.New("port must be greater than 0")
		}
		opts.port = port
		return nil
	}
}

func WithServerCA(serverCA []byte) NodeOption {
	return func(opts *NodeOptions) error {
		opts.serverCA = serverCA
		return nil
	}
}

func WithAPIKey(apiKey uuid.UUID) NodeOption {
	return func(opts *NodeOptions) error {
		opts.apiKey = apiKey
		return nil
	}
}

func WithExtra(extra map[string]interface{}) NodeOption {
	return func(opts *NodeOptions) error {
		opts.extra = extra
		return nil
	}
}

func WithLogChannelSize(size int) NodeOption {
	return func(opts *NodeOptions) error {
		if size <= 0 {
			opts.logChanSize = 1000
		} else {
			opts.logChanSize = size
		}
		return nil
	}
}

func New(address string, nodeProtocol NodeProtocol, options ...NodeOption) (BluePanelNode, error) {
	if address == "" {
		return nil, errors.New("address is empty")
	}

	opts := &NodeOptions{
		address:      address,
		nodeProtocol: nodeProtocol,
		extra:        make(map[string]interface{}),
	}

	for _, option := range options {
		if err := option(opts); err != nil {
			return nil, err
		}
	}

	var node BluePanelNode
	var err error
	switch nodeProtocol {
	case GRPC:
		node, err = rpc.New(opts.address, opts.port, opts.serverCA, opts.apiKey, opts.logChanSize, opts.extra)
	case REST:
		node, err = rest.New(opts.address, opts.port, opts.serverCA, opts.apiKey, opts.logChanSize, opts.extra)
	default:
		return nil, errors.New("unknown node protocol")
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}
