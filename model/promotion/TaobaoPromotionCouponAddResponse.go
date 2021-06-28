package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建店铺优惠券接口 APIResponse
taobao.promotion.coupon.add

创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
*/
type TaobaoPromotionCouponAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionCouponAddResponse `json:"promotion_coupon_add_response,omitempty"` 
    TaobaoPromotionCouponAddResponse
}

/* model for simplify = false
type TaobaoPromotionCouponAddResponse struct {

    // 优惠券的id
    
    CouponId   int64 `json:"coupon_id,omitempty"`
    

}
*/

type TaobaoPromotionCouponAddResponse struct {

    // 优惠券的id
    CouponId   int64 `json:"coupon_id,omitempty"`

}
