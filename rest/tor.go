package rest

import (
	"fmt"
	"net/http"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) ListTorLocations() (*common.TorLocationsResponse, error) {
	var out common.TorLocationsResponse
	if err := n.createRequest(n.client, http.MethodGet, "tor/locations", &common.Empty{}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (n *Node) GetTorLocation(id string) (*common.TorLocation, error) {
	var out common.TorLocation
	endpoint := fmt.Sprintf("tor/locations/%s/", id)
	if err := n.createRequest(n.client, http.MethodGet, endpoint, &common.Empty{}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (n *Node) CreateTorLocation(spec *common.TorLocationSpec) (*common.TorLocation, error) {
	var out common.TorLocation
	if err := n.createRequest(n.client, http.MethodPost, "tor/locations", spec, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (n *Node) UpdateTorLocation(spec *common.TorLocationSpec) (*common.TorLocation, error) {
	var out common.TorLocation
	endpoint := fmt.Sprintf("tor/locations/%s/", spec.GetId())
	if err := n.createRequest(n.client, http.MethodPut, endpoint, spec, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (n *Node) DeleteTorLocation(id string, purgeData bool) error {
	endpoint := fmt.Sprintf("tor/locations/%s/?purge_data=%t", id, purgeData)
	return n.createRequest(n.client, http.MethodDelete, endpoint, &common.Empty{}, &common.Empty{})
}

func (n *Node) torAction(action, id string) (*common.TorLocation, error) {
	var out common.TorLocation
	endpoint := fmt.Sprintf("tor/locations/%s/%s", id, action)
	if err := n.createRequest(n.client, http.MethodPost, endpoint, &common.Empty{}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (n *Node) EnableTorLocation(id string) (*common.TorLocation, error) {
	return n.torAction("enable", id)
}

func (n *Node) DisableTorLocation(id string) (*common.TorLocation, error) {
	return n.torAction("disable", id)
}

func (n *Node) RestartTorLocation(id string) (*common.TorLocation, error) {
	return n.torAction("restart", id)
}

func (n *Node) NewTorIdentity(id string) (*common.TorLocation, error) {
	return n.torAction("new-identity", id)
}

func (n *Node) GetTorHealth(id string) (*common.TorLocation, error) {
	return n.torAction("health", id)
}

func (n *Node) RepairTorLocation(id string) (*common.TorLocation, error) {
	return n.torAction("repair", id)
}

func (n *Node) TestTorLocation(id string) (*common.TorLocation, error) {
	return n.torAction("test", id)
}

func (n *Node) ForceReconcileTor() (*common.TorReconcileResponse, error) {
	var out common.TorReconcileResponse
	if err := n.createRequest(n.client, http.MethodPost, "tor/reconcile", &common.Empty{}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
