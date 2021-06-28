package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际机票大卖家Shopping推送 APIResponse
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。
*/
type TaobaoAlitripIeAgentShoppingPushAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentShoppingPushResponse
}

type TaobaoAlitripIeAgentShoppingPushResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_shopping_push_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ShoppingPushRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
