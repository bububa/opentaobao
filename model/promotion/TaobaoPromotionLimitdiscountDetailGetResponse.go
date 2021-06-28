package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
限时打折详情查询 APIResponse
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。
*/
type TaobaoPromotionLimitdiscountDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionLimitdiscountDetailGetResponse `json:"promotion_limitdiscount_detail_get_response,omitempty"` 
    TaobaoPromotionLimitdiscountDetailGetResponse
}

/* model for simplify = false
type TaobaoPromotionLimitdiscountDetailGetResponse struct {

    // 限时打折对应的商品详情列表。
    
    ItemDiscountDetailList  struct {
        LimitDiscountDetail  []LimitDiscountDetail `json:"limit_discount_detail,omitempty"`
    } `json:"item_discount_detail_list,omitempty"`
    

}
*/

type TaobaoPromotionLimitdiscountDetailGetResponse struct {

    // 限时打折对应的商品详情列表。
    ItemDiscountDetailList   []LimitDiscountDetail `json:"item_discount_detail_list,omitempty"`

}
