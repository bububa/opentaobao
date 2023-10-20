package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyRefuseAPIResponse 销售改签拒绝 API返回值
// alitrip.agent.flight.sell.modify.refuse
//
// 销售改签拒绝
type AlitripAgentFlightSellModifyRefuseAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellModifyRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellModifyRefuseAPIResponseModel).Reset()
}

// AlitripAgentFlightSellModifyRefuseAPIResponseModel is 销售改签拒绝 成功返回结果
type AlitripAgentFlightSellModifyRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellModifyRefuseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellModifyRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyRefuseAPIResponse)
	},
}

// GetAlitripAgentFlightSellModifyRefuseAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellModifyRefuseAPIResponse
func GetAlitripAgentFlightSellModifyRefuseAPIResponse() *AlitripAgentFlightSellModifyRefuseAPIResponse {
	return poolAlitripAgentFlightSellModifyRefuseAPIResponse.Get().(*AlitripAgentFlightSellModifyRefuseAPIResponse)
}

// ReleaseAlitripAgentFlightSellModifyRefuseAPIResponse 将 AlitripAgentFlightSellModifyRefuseAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellModifyRefuseAPIResponse(v *AlitripAgentFlightSellModifyRefuseAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellModifyRefuseAPIResponse.Put(v)
}
