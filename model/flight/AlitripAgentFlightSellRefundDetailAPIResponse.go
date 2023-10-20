package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellRefundDetailAPIResponse 销售退票单详情 API返回值
// alitrip.agent.flight.sell.refund.detail
//
// 销售退票单详情
type AlitripAgentFlightSellRefundDetailAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellRefundDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellRefundDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellRefundDetailAPIResponseModel).Reset()
}

// AlitripAgentFlightSellRefundDetailAPIResponseModel is 销售退票单详情 成功返回结果
type AlitripAgentFlightSellRefundDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellRefundDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellRefundDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellRefundDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellRefundDetailAPIResponse)
	},
}

// GetAlitripAgentFlightSellRefundDetailAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellRefundDetailAPIResponse
func GetAlitripAgentFlightSellRefundDetailAPIResponse() *AlitripAgentFlightSellRefundDetailAPIResponse {
	return poolAlitripAgentFlightSellRefundDetailAPIResponse.Get().(*AlitripAgentFlightSellRefundDetailAPIResponse)
}

// ReleaseAlitripAgentFlightSellRefundDetailAPIResponse 将 AlitripAgentFlightSellRefundDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellRefundDetailAPIResponse(v *AlitripAgentFlightSellRefundDetailAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellRefundDetailAPIResponse.Put(v)
}
