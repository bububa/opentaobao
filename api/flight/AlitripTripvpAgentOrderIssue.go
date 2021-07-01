package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

/* AlitripTripvpAgentOrderIssue
廉航辅营正向订单出货接口
alitrip.tripvp.agent.order.issue

廉航辅营正向订单出货接口 */
func AlitripTripvpAgentOrderIssue(clt *core.SDKClient, req *flight.AlitripTripvpAgentOrderIssueAPIRequest, session string) (*flight.AlitripTripvpAgentOrderIssueAPIResponse, error) {
	var resp flight.AlitripTripvpAgentOrderIssueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
