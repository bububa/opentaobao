package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyListAPIResponse 销售改签单列表 API返回值
// alitrip.agent.flight.sell.modify.list
//
// 销售改签单列表
type AlitripAgentFlightSellModifyListAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellModifyListAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellModifyListAPIResponseModel).Reset()
}

// AlitripAgentFlightSellModifyListAPIResponseModel is 销售改签单列表 成功返回结果
type AlitripAgentFlightSellModifyListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellModifyListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellModifyListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyListAPIResponse)
	},
}

// GetAlitripAgentFlightSellModifyListAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellModifyListAPIResponse
func GetAlitripAgentFlightSellModifyListAPIResponse() *AlitripAgentFlightSellModifyListAPIResponse {
	return poolAlitripAgentFlightSellModifyListAPIResponse.Get().(*AlitripAgentFlightSellModifyListAPIResponse)
}

// ReleaseAlitripAgentFlightSellModifyListAPIResponse 将 AlitripAgentFlightSellModifyListAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellModifyListAPIResponse(v *AlitripAgentFlightSellModifyListAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellModifyListAPIResponse.Put(v)
}
