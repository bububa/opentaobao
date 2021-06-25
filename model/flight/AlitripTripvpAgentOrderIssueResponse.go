package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单出货接口 APIResponse
alitrip.tripvp.agent.order.issue

廉航辅营正向订单出货接口
*/
type AlitripTripvpAgentOrderIssueAPIResponse struct {
    model.CommonResponse
    Response *AlitripTripvpAgentOrderIssueResponse `json:"alitrip_tripvp_agent_order_issue_response,omitempty"`
}

type AlitripTripvpAgentOrderIssueResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
