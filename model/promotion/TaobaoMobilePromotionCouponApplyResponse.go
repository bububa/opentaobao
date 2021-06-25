package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取(手淘专用) APIResponse
taobao.mobile.promotion.coupon.apply

优惠券领取
*/
type TaobaoMobilePromotionCouponApplyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMobilePromotionCouponApplyResponse `json:"taobao_mobile_promotion_coupon_apply_response,omitempty"`
}

type TaobaoMobilePromotionCouponApplyResponse struct {

    // 优惠券领取结果
    CouponApplyResult   *CouponApplyResult `json:"coupon_apply_result,omitempty"`

}
