package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyDetailAPIResponse 销售改签详情 API返回值
// alitrip.agent.flight.sell.modify.detail
//
// 销售改签详情
type AlitripAgentFlightSellModifyDetailAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellModifyDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellModifyDetailAPIResponseModel).Reset()
}

// AlitripAgentFlightSellModifyDetailAPIResponseModel is 销售改签详情 成功返回结果
type AlitripAgentFlightSellModifyDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellModifyDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellModifyDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyDetailAPIResponse)
	},
}

// GetAlitripAgentFlightSellModifyDetailAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellModifyDetailAPIResponse
func GetAlitripAgentFlightSellModifyDetailAPIResponse() *AlitripAgentFlightSellModifyDetailAPIResponse {
	return poolAlitripAgentFlightSellModifyDetailAPIResponse.Get().(*AlitripAgentFlightSellModifyDetailAPIResponse)
}

// ReleaseAlitripAgentFlightSellModifyDetailAPIResponse 将 AlitripAgentFlightSellModifyDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellModifyDetailAPIResponse(v *AlitripAgentFlightSellModifyDetailAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellModifyDetailAPIResponse.Put(v)
}
