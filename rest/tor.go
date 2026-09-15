package rest

import (
	"net/http"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) ListTorLocations() (*common.TorLocationsResponse, error) {
	var out common.TorLocationsResponse
	if err := n.createRequest(n.client, http.MethodGet, "tor/locations", &common.Empty{}, &out); err != nil { return nil, err }
	return &out, nil
}

func (n *Node) GetTorLocation(id string) (*common.TorLocation, error) {
	var out common.TorLocation
	if err := n.createRequest(n.client, http.MethodPost, "tor/location/get", &common.TorLocationIDRequest{Id: id}, &out); err != nil { return nil, err }
	return &out, nil
}

func (n *Node) CreateTorLocation(spec *common.TorLocationSpec) (*common.TorLocation, error) {
	var out common.TorLocation
	if err := n.createRequest(n.client, http.MethodPost, "tor/location", spec, &out); err != nil { return nil, err }
	return &out, nil
}

func (n *Node) UpdateTorLocation(spec *common.TorLocationSpec) (*common.TorLocation, error) {
	var out common.TorLocation
	if err := n.createRequest(n.client, http.MethodPut, "tor/location", spec, &out); err != nil { return nil, err }
	return &out, nil
}

func (n *Node) DeleteTorLocation(id string, purgeData bool) error {
	return n.createRequest(n.client, http.MethodPost, "tor/location/delete", &common.DeleteTorLocationRequest{Id: id, PurgeData: purgeData}, &common.Empty{})
}

func (n *Node) torAction(endpoint, id string) (*common.TorLocation, error) {
	var out common.TorLocation
	if err := n.createRequest(n.client, http.MethodPost, endpoint, &common.TorLocationIDRequest{Id: id}, &out); err != nil { return nil, err }
	return &out, nil
}

func (n *Node) EnableTorLocation(id string) (*common.TorLocation, error) { return n.torAction("tor/location/enable", id) }
func (n *Node) DisableTorLocation(id string) (*common.TorLocation, error) { return n.torAction("tor/location/disable", id) }
func (n *Node) RestartTorLocation(id string) (*common.TorLocation, error) { return n.torAction("tor/location/restart", id) }
func (n *Node) NewTorIdentity(id string) (*common.TorLocation, error) { return n.torAction("tor/location/new-identity", id) }
func (n *Node) GetTorHealth(id string) (*common.TorLocation, error) { return n.torAction("tor/location/health", id) }
func (n *Node) RepairTorLocation(id string) (*common.TorLocation, error) { return n.torAction("tor/location/repair", id) }
func (n *Node) TestTorLocation(id string) (*common.TorLocation, error) { return n.torAction("tor/location/test", id) }

func (n *Node) ForceReconcileTor() (*common.TorReconcileResponse, error) {
	var out common.TorReconcileResponse
	if err := n.createRequest(n.client, http.MethodPost, "tor/reconcile", &common.Empty{}, &out); err != nil { return nil, err }
	return &out, nil
}
