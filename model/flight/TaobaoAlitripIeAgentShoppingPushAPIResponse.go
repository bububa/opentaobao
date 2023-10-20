package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentShoppingPushAPIResponse 国际机票大卖家Shopping推送 API返回值
// taobao.alitrip.ie.agent.shopping.push
//
// 用于国际机票大卖家主动推送Shopping结果更新缓存报价。
type TaobaoAlitripIeAgentShoppingPushAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentShoppingPushAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentShoppingPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentShoppingPushAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentShoppingPushAPIResponseModel is 国际机票大卖家Shopping推送 成功返回结果
type TaobaoAlitripIeAgentShoppingPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_shopping_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ShoppingPushRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentShoppingPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentShoppingPushAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentShoppingPushAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentShoppingPushAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentShoppingPushAPIResponse
func GetTaobaoAlitripIeAgentShoppingPushAPIResponse() *TaobaoAlitripIeAgentShoppingPushAPIResponse {
	return poolTaobaoAlitripIeAgentShoppingPushAPIResponse.Get().(*TaobaoAlitripIeAgentShoppingPushAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentShoppingPushAPIResponse 将 TaobaoAlitripIeAgentShoppingPushAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentShoppingPushAPIResponse(v *TaobaoAlitripIeAgentShoppingPushAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentShoppingPushAPIResponse.Put(v)
}
