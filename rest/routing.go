package rest

import (
	"net/http"

	"github.com/hazhanhasani/node_bridge/common"
)

func (n *Node) ListRoutingRules() (*common.RoutingRulesResponse, error) {
	var response common.RoutingRulesResponse
	if err := n.createRequest(n.client, http.MethodGet, "routing/rules", &common.Empty{}, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (n *Node) GetBalancerInfo(tag string) (*common.BalancerInfoResponse, error) {
	var response common.BalancerInfoResponse
	if err := n.createRequest(n.client, http.MethodPost, "routing/balancer", &common.BalancerInfoRequest{Tag: tag}, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (n *Node) TestRoute(request *common.TestRouteRequest) (*common.RouteResult, error) {
	if request == nil {
		request = &common.TestRouteRequest{}
	}
	var response common.RouteResult
	if err := n.createRequest(n.client, http.MethodPost, "routing/test", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (n *Node) AddRoutingRule(rule string, shouldReset bool) error {
	return n.createRequest(n.client, http.MethodPut, "routing/rules", &common.AddRoutingRuleRequest{Rule: rule, ShouldReset: shouldReset}, &common.Empty{})
}

func (n *Node) RemoveRoutingRule(ruleTag string) error {
	return n.createRequest(n.client, http.MethodDelete, "routing/rules", &common.RemoveRoutingRuleRequest{RuleTag: ruleTag}, &common.Empty{})
}

func (n *Node) OverrideBalancerTarget(balancerTag, target string) error {
	return n.createRequest(n.client, http.MethodPut, "routing/balancer/override", &common.OverrideBalancerTargetRequest{BalancerTag: balancerTag, Target: target}, &common.Empty{})
}
