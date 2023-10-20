package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellRefundRefuseAPIResponse 销售退票拒绝 API返回值
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
type AlitripAgentFlightSellRefundRefuseAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellRefundRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellRefundRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellRefundRefuseAPIResponseModel).Reset()
}

// AlitripAgentFlightSellRefundRefuseAPIResponseModel is 销售退票拒绝 成功返回结果
type AlitripAgentFlightSellRefundRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellRefundRefuseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellRefundRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellRefundRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellRefundRefuseAPIResponse)
	},
}

// GetAlitripAgentFlightSellRefundRefuseAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellRefundRefuseAPIResponse
func GetAlitripAgentFlightSellRefundRefuseAPIResponse() *AlitripAgentFlightSellRefundRefuseAPIResponse {
	return poolAlitripAgentFlightSellRefundRefuseAPIResponse.Get().(*AlitripAgentFlightSellRefundRefuseAPIResponse)
}

// ReleaseAlitripAgentFlightSellRefundRefuseAPIResponse 将 AlitripAgentFlightSellRefundRefuseAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellRefundRefuseAPIResponse(v *AlitripAgentFlightSellRefundRefuseAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellRefundRefuseAPIResponse.Put(v)
}
