package flight

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripTripvpAgentOrderIssueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTripvpAgentOrderIssueAPIResponseModel).Reset()
}

// AlitripTripvpAgentOrderIssueAPIResponseModel is 廉航辅营正向订单出货接口 成功返回结果
type AlitripTripvpAgentOrderIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tripvp_agent_order_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTripvpAgentOrderIssueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlitripTripvpAgentOrderIssueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTripvpAgentOrderIssueAPIResponse)
	},
}

// GetAlitripTripvpAgentOrderIssueAPIResponse 从 sync.Pool 获取 AlitripTripvpAgentOrderIssueAPIResponse
func GetAlitripTripvpAgentOrderIssueAPIResponse() *AlitripTripvpAgentOrderIssueAPIResponse {
	return poolAlitripTripvpAgentOrderIssueAPIResponse.Get().(*AlitripTripvpAgentOrderIssueAPIResponse)
}

// ReleaseAlitripTripvpAgentOrderIssueAPIResponse 将 AlitripTripvpAgentOrderIssueAPIResponse 保存到 sync.Pool
func ReleaseAlitripTripvpAgentOrderIssueAPIResponse(v *AlitripTripvpAgentOrderIssueAPIResponse) {
	v.Reset()
	poolAlitripTripvpAgentOrderIssueAPIResponse.Put(v)
}
