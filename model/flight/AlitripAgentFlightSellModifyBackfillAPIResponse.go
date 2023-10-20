package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyBackfillAPIResponse 销售改签回填 API返回值
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
type AlitripAgentFlightSellModifyBackfillAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellModifyBackfillAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyBackfillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellModifyBackfillAPIResponseModel).Reset()
}

// AlitripAgentFlightSellModifyBackfillAPIResponseModel is 销售改签回填 成功返回结果
type AlitripAgentFlightSellModifyBackfillAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_backfill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellModifyBackfillResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyBackfillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellModifyBackfillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyBackfillAPIResponse)
	},
}

// GetAlitripAgentFlightSellModifyBackfillAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellModifyBackfillAPIResponse
func GetAlitripAgentFlightSellModifyBackfillAPIResponse() *AlitripAgentFlightSellModifyBackfillAPIResponse {
	return poolAlitripAgentFlightSellModifyBackfillAPIResponse.Get().(*AlitripAgentFlightSellModifyBackfillAPIResponse)
}

// ReleaseAlitripAgentFlightSellModifyBackfillAPIResponse 将 AlitripAgentFlightSellModifyBackfillAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellModifyBackfillAPIResponse(v *AlitripAgentFlightSellModifyBackfillAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellModifyBackfillAPIResponse.Put(v)
}
