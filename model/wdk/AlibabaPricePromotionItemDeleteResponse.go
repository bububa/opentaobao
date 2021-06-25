package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量删除档期 APIResponse
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品
*/
type AlibabaPricePromotionItemDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaPricePromotionItemDeleteResponse `json:"alibaba_price_promotion_item_delete_response,omitempty"`
}

type AlibabaPricePromotionItemDeleteResponse struct {

    // result
    Result   *AlibabaPricePromotionItemDeleteResult `json:"result,omitempty"`

}
