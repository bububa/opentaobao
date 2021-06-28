package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家最低折扣 APIResponse
taobao.ump.promotion.global.discount.get

提供卖家最低折扣查询功能
*/
type TaobaoUmpPromotionGlobalDiscountGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpPromotionGlobalDiscountGetResponse `json:"ump_promotion_global_discount_get_response,omitempty"` 
    TaobaoUmpPromotionGlobalDiscountGetResponse
}

/* model for simplify = false
type TaobaoUmpPromotionGlobalDiscountGetResponse struct {

    // 结果对象
    
    Result  *struct {
        TaobaoUmpPromotionGlobalDiscountGetResult  *TaobaoUmpPromotionGlobalDiscountGetResult `json:"taobao_ump_promotion_global_discount_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoUmpPromotionGlobalDiscountGetResponse struct {

    // 结果对象
    Result   *TaobaoUmpPromotionGlobalDiscountGetResult `json:"result,omitempty"`

}
