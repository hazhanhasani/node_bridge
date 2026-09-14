package rpc

import (
	"context"
	"time"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) ListRoutingRules() (*common.RoutingRulesResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	return n.client.ListRoutingRules(ctx, &common.Empty{})
}

func (n *Node) GetBalancerInfo(tag string) (*common.BalancerInfoResponse, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	return n.client.GetBalancerInfo(ctx, &common.BalancerInfoRequest{Tag: tag})
}

func (n *Node) TestRoute(request *common.TestRouteRequest) (*common.RouteResult, error) {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	if request == nil {
		request = &common.TestRouteRequest{}
	}
	return n.client.TestRoute(ctx, request)
}

func (n *Node) AddRoutingRule(rule string, shouldReset bool) error {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	_, err := n.client.AddRoutingRule(ctx, &common.AddRoutingRuleRequest{Rule: rule, ShouldReset: shouldReset})
	return err
}

func (n *Node) RemoveRoutingRule(ruleTag string) error {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	_, err := n.client.RemoveRoutingRule(ctx, &common.RemoveRoutingRuleRequest{RuleTag: ruleTag})
	return err
}

func (n *Node) OverrideBalancerTarget(balancerTag, target string) error {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	_, err := n.client.OverrideBalancerTarget(ctx, &common.OverrideBalancerTargetRequest{BalancerTag: balancerTag, Target: target})
	return err
}
