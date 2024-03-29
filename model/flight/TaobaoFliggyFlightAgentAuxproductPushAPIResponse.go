package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyFlightAgentAuxproductPushAPIResponse 飞猪机票辅营商品投放 API返回值
// taobao.fliggy.flight.agent.auxproduct.push
//
// 廉航辅营产品投放接口
type TaobaoFliggyFlightAgentAuxproductPushAPIResponse struct {
	model.CommonResponse
	TaobaoFliggyFlightAgentAuxproductPushAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFliggyFlightAgentAuxproductPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFliggyFlightAgentAuxproductPushAPIResponseModel).Reset()
}

// TaobaoFliggyFlightAgentAuxproductPushAPIResponseModel is 飞猪机票辅营商品投放 成功返回结果
type TaobaoFliggyFlightAgentAuxproductPushAPIResponseModel struct {
	XMLName xml.Name `xml:"fliggy_flight_agent_auxproduct_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作日志id，商家可通过该id在后台查看本次操作的具体结果
	TracerId string `json:"tracer_id,omitempty" xml:"tracer_id,omitempty"`
	// 备注
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFliggyFlightAgentAuxproductPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TracerId = ""
	m.Message = ""
	m.IsSuccess = false
}

var poolTaobaoFliggyFlightAgentAuxproductPushAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFliggyFlightAgentAuxproductPushAPIResponse)
	},
}

// GetTaobaoFliggyFlightAgentAuxproductPushAPIResponse 从 sync.Pool 获取 TaobaoFliggyFlightAgentAuxproductPushAPIResponse
func GetTaobaoFliggyFlightAgentAuxproductPushAPIResponse() *TaobaoFliggyFlightAgentAuxproductPushAPIResponse {
	return poolTaobaoFliggyFlightAgentAuxproductPushAPIResponse.Get().(*TaobaoFliggyFlightAgentAuxproductPushAPIResponse)
}

// ReleaseTaobaoFliggyFlightAgentAuxproductPushAPIResponse 将 TaobaoFliggyFlightAgentAuxproductPushAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFliggyFlightAgentAuxproductPushAPIResponse(v *TaobaoFliggyFlightAgentAuxproductPushAPIResponse) {
	v.Reset()
	poolTaobaoFliggyFlightAgentAuxproductPushAPIResponse.Put(v)
}
