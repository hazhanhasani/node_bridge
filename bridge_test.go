package bluepanel_node_bridge

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hazhanhasani/node_bridge/common"
	"github.com/hazhanhasani/node_bridge/tools"
)

const defaultE2EPort = 2096

// loadE2EConfig keeps live-node tests opt-in. Unit/CI test runs must not depend
// on a specific server, certificate, API key, or local config file.
func loadE2EConfig(t *testing.T) (string, string, uint64, []NodeOption) {
	t.Helper()

	if os.Getenv("BLUEPANEL_NODE_E2E") != "1" {
		t.Skip("set BLUEPANEL_NODE_E2E=1 to run live BluePanel Node integration tests")
	}

	nodeAddr := os.Getenv("BLUEPANEL_NODE_ADDR")
	serverCAPath := os.Getenv("BLUEPANEL_NODE_CA")
	apiKeyValue := os.Getenv("BLUEPANEL_NODE_API_KEY")
	configPath := os.Getenv("BLUEPANEL_NODE_CONFIG")
	if nodeAddr == "" || serverCAPath == "" || apiKeyValue == "" || configPath == "" {
		t.Fatal("BLUEPANEL_NODE_ADDR, BLUEPANEL_NODE_CA, BLUEPANEL_NODE_API_KEY and BLUEPANEL_NODE_CONFIG are required for E2E tests")
	}

	port := defaultE2EPort
	if value := os.Getenv("BLUEPANEL_NODE_PORT"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 || parsed > 65535 {
			t.Fatalf("invalid BLUEPANEL_NODE_PORT %q", value)
		}
		port = parsed
	}

	serverCA, err := os.ReadFile(serverCAPath)
	if err != nil {
		t.Fatalf("read BluePanel Node CA: %v", err)
	}

	apiKey, err := uuid.Parse(apiKeyValue)
	if err != nil {
		t.Fatalf("parse BLUEPANEL_NODE_API_KEY: %v", err)
	}

	configFile, err := tools.ReadFileAsString(configPath)
	if err != nil {
		t.Fatalf("read BLUEPANEL_NODE_CONFIG: %v", err)
	}

	return nodeAddr, configFile, 60, []NodeOption{
		WithPort(port),
		WithAPIKey(apiKey),
		WithServerCA(serverCA),
		WithLogChannelSize(100),
	}
}

func testUser() *common.User {
	return common.CreateUser(
		"test_user",
		common.CreateProxies(
			common.CreateVmess(uuid.New().String()),
			common.CreateVless(uuid.New().String(), ""),
			common.CreateTrojan("random data"),
			common.CreateShadowsocks("random", "aes-256-gcm"),
			nil,
			nil,
		),
		[]string{},
	)
}

func TestGrpcNodeE2E(t *testing.T) {
	nodeAddr, configFile, keepAlive, opts := loadE2EConfig(t)
	node, err := New(nodeAddr, GRPC, opts...)
	if err != nil {
		t.Fatal(err)
	}

	if err = node.Start(configFile, common.BackendType_XRAY, nil, keepAlive); err != nil {
		t.Fatal(err)
	}
	defer node.Stop()

	info, err := node.Info()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("node info: %+v", info)

	node.UpdateUsers([]*common.User{testUser()})

	_, err = node.GetUserOnlineIpList("does-not-exist@example.com")
	st, _ := status.FromError(err)
	if st.Code() != codes.NotFound {
		t.Fatalf("expected NotFound for missing user, got %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logChan, err := node.StreamLogs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for entry := range logChan {
			if entry.Err != nil {
				t.Logf("log stream ended: %v", entry.Err)
				return
			}
			t.Log(entry.Line)
		}
	}()

	time.Sleep(2 * time.Second)
}

func TestRestNodeE2E(t *testing.T) {
	nodeAddr, configFile, keepAlive, opts := loadE2EConfig(t)
	node, err := New(nodeAddr, REST, opts...)
	if err != nil {
		t.Fatal(err)
	}

	if err = node.Start(configFile, common.BackendType_XRAY, nil, keepAlive); err != nil {
		t.Fatal(err)
	}
	defer node.Stop()

	info, err := node.Info()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("node info: %+v", info)

	node.UpdateUsers([]*common.User{testUser()})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logChan, err := node.StreamLogs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for entry := range logChan {
			if entry.Err != nil {
				t.Logf("log stream ended: %v", entry.Err)
				return
			}
			t.Log(entry.Line)
		}
	}()

	time.Sleep(3 * time.Second)
	stats, err := node.GetStats(true, "", common.StatType_Outbounds)
	if err != nil {
		t.Fatal(err)
	}
	for _, stat := range stats.GetStats() {
		t.Logf("stat: %+v", stat)
	}
}
