package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单出货接口 APIResponse
alitrip.tripvp.agent.order.issue

廉航辅营正向订单出货接口
*/
type AlitripTripvpAgentOrderIssueAPIResponse struct {
    model.CommonResponse
    AlitripTripvpAgentOrderIssueResponse
}

type AlitripTripvpAgentOrderIssueResponse struct {
    XMLName xml.Name `xml:"alitrip_tripvp_agent_order_issue_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
