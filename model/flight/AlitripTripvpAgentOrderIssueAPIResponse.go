package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTripvpAgentOrderIssueAPIResponse 廉航辅营正向订单出货接口 API返回值
// alitrip.tripvp.agent.order.issue
//
// 廉航辅营正向订单出货接口
type AlitripTripvpAgentOrderIssueAPIResponse struct {
	model.CommonResponse
	AlitripTripvpAgentOrderIssueAPIResponseModel
}

// AlitripTripvpAgentOrderIssueAPIResponseModel is 廉航辅营正向订单出货接口 成功返回结果
type AlitripTripvpAgentOrderIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tripvp_agent_order_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
