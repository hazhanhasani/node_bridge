package rpc

import (
	"context"
	"time"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) torContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(n.ctx, 30*time.Second)
}

func (n *Node) ListTorLocations() (*common.TorLocationsResponse, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.ListTorLocations(ctx, &common.Empty{})
}

func (n *Node) GetTorLocation(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.GetTorLocation(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) CreateTorLocation(spec *common.TorLocationSpec) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.CreateTorLocation(ctx, spec)
}

func (n *Node) UpdateTorLocation(spec *common.TorLocationSpec) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.UpdateTorLocation(ctx, spec)
}

func (n *Node) DeleteTorLocation(id string, purgeData bool) error {
	ctx, cancel := n.torContext()
	defer cancel()
	_, err := n.client.DeleteTorLocation(ctx, &common.DeleteTorLocationRequest{Id: id, PurgeData: purgeData})
	return err
}

func (n *Node) EnableTorLocation(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.EnableTorLocation(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) DisableTorLocation(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.DisableTorLocation(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) RestartTorLocation(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.RestartTorLocation(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) NewTorIdentity(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.NewTorIdentity(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) GetTorHealth(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.GetTorHealth(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) RepairTorLocation(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.RepairTorLocation(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) TestTorLocation(id string) (*common.TorLocation, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.TestTorLocation(ctx, &common.TorLocationIDRequest{Id: id})
}

func (n *Node) ForceReconcileTor() (*common.TorReconcileResponse, error) {
	ctx, cancel := n.torContext()
	defer cancel()
	return n.client.ForceReconcileTor(ctx, &common.Empty{})
}
