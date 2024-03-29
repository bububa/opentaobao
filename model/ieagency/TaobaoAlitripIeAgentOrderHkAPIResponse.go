package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentOrderHkAPIResponse 【国际机票】手工预定回填PNR API返回值
// taobao.alitrip.ie.agent.order.hk
//
// 代理商通过手工预定PNR，并回填。
type TaobaoAlitripIeAgentOrderHkAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentOrderHkAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentOrderHkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentOrderHkAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentOrderHkAPIResponseModel is 【国际机票】手工预定回填PNR 成功返回结果
type TaobaoAlitripIeAgentOrderHkAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_order_hk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否回填成功true：成功 false:失败
	HkSuccess bool `json:"hk_success,omitempty" xml:"hk_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentOrderHkAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HkSuccess = false
}

var poolTaobaoAlitripIeAgentOrderHkAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentOrderHkAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentOrderHkAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentOrderHkAPIResponse
func GetTaobaoAlitripIeAgentOrderHkAPIResponse() *TaobaoAlitripIeAgentOrderHkAPIResponse {
	return poolTaobaoAlitripIeAgentOrderHkAPIResponse.Get().(*TaobaoAlitripIeAgentOrderHkAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentOrderHkAPIResponse 将 TaobaoAlitripIeAgentOrderHkAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentOrderHkAPIResponse(v *TaobaoAlitripIeAgentOrderHkAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentOrderHkAPIResponse.Put(v)
}
