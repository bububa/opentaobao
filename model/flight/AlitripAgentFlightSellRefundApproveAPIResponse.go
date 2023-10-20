package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellRefundApproveAPIResponse 销售退票确认 API返回值
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
type AlitripAgentFlightSellRefundApproveAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellRefundApproveAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellRefundApproveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellRefundApproveAPIResponseModel).Reset()
}

// AlitripAgentFlightSellRefundApproveAPIResponseModel is 销售退票确认 成功返回结果
type AlitripAgentFlightSellRefundApproveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_approve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellRefundApproveResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellRefundApproveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellRefundApproveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellRefundApproveAPIResponse)
	},
}

// GetAlitripAgentFlightSellRefundApproveAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellRefundApproveAPIResponse
func GetAlitripAgentFlightSellRefundApproveAPIResponse() *AlitripAgentFlightSellRefundApproveAPIResponse {
	return poolAlitripAgentFlightSellRefundApproveAPIResponse.Get().(*AlitripAgentFlightSellRefundApproveAPIResponse)
}

// ReleaseAlitripAgentFlightSellRefundApproveAPIResponse 将 AlitripAgentFlightSellRefundApproveAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellRefundApproveAPIResponse(v *AlitripAgentFlightSellRefundApproveAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellRefundApproveAPIResponse.Put(v)
}
