package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse 飞猪机票辅营商品删除 API返回值
// taobao.fliggy.flight.agent.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoFliggyFlightAgentAuxproductDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFliggyFlightAgentAuxproductDeleteAPIResponseModel).Reset()
}

// TaobaoFliggyFlightAgentAuxproductDeleteAPIResponseModel is 飞猪机票辅营商品删除 成功返回结果
type TaobaoFliggyFlightAgentAuxproductDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"fliggy_flight_agent_auxproduct_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DelAuxProductsRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFliggyFlightAgentAuxproductDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse)
	},
}

// GetTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse 从 sync.Pool 获取 TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse
func GetTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse() *TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse {
	return poolTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse.Get().(*TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse)
}

// ReleaseTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse 将 TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse(v *TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse) {
	v.Reset()
	poolTaobaoFliggyFlightAgentAuxproductDeleteAPIResponse.Put(v)
}
