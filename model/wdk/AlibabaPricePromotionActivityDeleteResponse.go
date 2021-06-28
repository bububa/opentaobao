package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除档期活动 APIResponse
alibaba.price.promotion.activity.delete

删除盒马帮档期活动
*/
type AlibabaPricePromotionActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaPricePromotionActivityDeleteResponse `json:"alibaba_price_promotion_activity_delete_response,omitempty"` 
    AlibabaPricePromotionActivityDeleteResponse
}

/* model for simplify = false
type AlibabaPricePromotionActivityDeleteResponse struct {

    // result
    
    Result  *struct {
        AlibabaPricePromotionActivityDeleteResult  *AlibabaPricePromotionActivityDeleteResult `json:"alibaba_price_promotion_activity_delete_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaPricePromotionActivityDeleteResponse struct {

    // result
    Result   *AlibabaPricePromotionActivityDeleteResult `json:"result,omitempty"`

}
