package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家优惠券 APIResponse
taobao.promotion.coupons.get

查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量
*/
type TaobaoPromotionCouponsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionCouponsGetResponse `json:"promotion_coupons_get_response,omitempty"` 
    TaobaoPromotionCouponsGetResponse
}

/* model for simplify = false
type TaobaoPromotionCouponsGetResponse struct {

    // 查询的总数量
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 优惠券列表
    
    Coupons  struct {
        Coupon  []Coupon `json:"coupon,omitempty"`
    } `json:"coupons,omitempty"`
    

}
*/

type TaobaoPromotionCouponsGetResponse struct {

    // 查询的总数量
    TotalResults   int64 `json:"total_results,omitempty"`

    // 优惠券列表
    Coupons   []Coupon `json:"coupons,omitempty"`

}
