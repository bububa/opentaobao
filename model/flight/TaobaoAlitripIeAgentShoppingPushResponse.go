package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际机票大卖家Shopping推送 API返回值 
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。
*/
type TaobaoAlitripIeAgentShoppingPushAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentShoppingPushResponse
}

// 国际机票大卖家Shopping推送 成功返回结果
type TaobaoAlitripIeAgentShoppingPushResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_shopping_push_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ShoppingPushRs `json:"result,omitempty" xml:"result,omitempty"`
}
