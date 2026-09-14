package rest

import (
	"net/http"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) GetSystemStats() (*common.SystemStatsResponse, error) {
	var stats common.SystemStatsResponse
	if err := n.createRequest(n.client, http.MethodGet, "stats/system", &common.Empty{}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (n *Node) GetBackendStats() (*common.BackendStatsResponse, error) {
	var stats common.BackendStatsResponse
	if err := n.createRequest(n.client, http.MethodGet, "stats/backend", &common.Empty{}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (n *Node) GetStats(reset bool, name string, statType common.StatType) (*common.StatResponse, error) {
	var stats common.StatResponse
	if err := n.createRequest(n.client, http.MethodGet, "stats/", &common.StatRequest{Reset_: reset, Name: name, Type: statType}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (n *Node) GetOutboundsLatency(name string) (*common.LatencyResponse, error) {
	var latency common.LatencyResponse
	if err := n.createRequest(n.client, http.MethodGet, "stats/latency", &common.LatencyRequest{Name: name}, &latency); err != nil {
		return nil, err
	}
	return &latency, nil
}

func (n *Node) GetUserOnlineStat(email string) (*common.OnlineStatResponse, error) {
	var stats common.OnlineStatResponse
	if err := n.createRequest(n.client, http.MethodGet, "stats/user/online", &common.StatRequest{Name: email}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (n *Node) GetUserOnlineIpList(email string) (*common.StatsOnlineIpListResponse, error) {
	var stats common.StatsOnlineIpListResponse
	if err := n.createRequest(n.client, http.MethodGet, "stats/user/online_ip", &common.StatRequest{Name: email}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}
