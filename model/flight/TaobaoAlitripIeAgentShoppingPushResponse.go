package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
国际机票大卖家Shopping推送 APIResponse
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。
*/
type TaobaoAlitripIeAgentShoppingPushAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripIeAgentShoppingPushResponse `json:"alitrip_ie_agent_shopping_push_response,omitempty"` 
    TaobaoAlitripIeAgentShoppingPushResponse
}

/* model for simplify = false
type TaobaoAlitripIeAgentShoppingPushResponse struct {

    // result
    
    Result  *struct {
        ShoppingPushRs  *ShoppingPushRs `json:"shopping_push_rs,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAlitripIeAgentShoppingPushResponse struct {

    // result
    Result   *ShoppingPushRs `json:"result,omitempty"`

}
