package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增档期商品 APIResponse
alibaba.price.promotion.item.add

批量新增档期活动商品
*/
type AlibabaPricePromotionItemAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaPricePromotionItemAddResponse `json:"alibaba_price_promotion_item_add_response,omitempty"` 
    AlibabaPricePromotionItemAddResponse
}

/* model for simplify = false
type AlibabaPricePromotionItemAddResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaPricePromotionItemAddResult  *AlibabaPricePromotionItemAddResult `json:"alibaba_price_promotion_item_add_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaPricePromotionItemAddResponse struct {

    // 返回结果
    Result   *AlibabaPricePromotionItemAddResult `json:"result,omitempty"`

}
