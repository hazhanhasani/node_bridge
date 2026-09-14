package rpc

import (
	"context"
	"time"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) GetSystemStats() (*common.SystemStatsResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	defer cancel()
	return n.client.GetSystemStats(ctx, nil)
}

func (n *Node) GetBackendStats() (*common.BackendStatsResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	defer cancel()
	return n.client.GetBackendStats(ctx, nil)
}

func (n *Node) GetStats(reset bool, name string, statType common.StatType) (*common.StatResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	defer cancel()
	return n.client.GetStats(ctx, &common.StatRequest{Reset_: reset, Name: name, Type: statType})
}

func (n *Node) GetOutboundsLatency(name string) (*common.LatencyResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	return n.client.GetOutboundsLatency(ctx, &common.LatencyRequest{Name: name})
}

func (n *Node) GetUserOnlineStat(email string) (*common.OnlineStatResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	defer cancel()
	return n.client.GetUserOnlineStats(ctx, &common.StatRequest{Name: email})
}

func (n *Node) GetUserOnlineIpList(email string) (*common.StatsOnlineIpListResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	defer cancel()
	return n.client.GetUserOnlineIpListStats(ctx, &common.StatRequest{Name: email})
}
