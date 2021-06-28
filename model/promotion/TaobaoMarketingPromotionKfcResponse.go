package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向优惠活动名称与描述违禁词检查 APIResponse
taobao.marketing.promotion.kfc

活动名称与描述违禁词检查
*/
type TaobaoMarketingPromotionKfcAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMarketingPromotionKfcResponse `json:"marketing_promotion_kfc_response,omitempty"` 
    TaobaoMarketingPromotionKfcResponse
}

/* model for simplify = false
type TaobaoMarketingPromotionKfcResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoMarketingPromotionKfcResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
