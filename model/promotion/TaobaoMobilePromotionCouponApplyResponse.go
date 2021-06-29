package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取(手淘专用) API返回值 
taobao.mobile.promotion.coupon.apply

优惠券领取
*/
type TaobaoMobilePromotionCouponApplyAPIResponse struct {
    model.CommonResponse
    TaobaoMobilePromotionCouponApplyResponse
}

// 优惠券领取(手淘专用) 成功返回结果
type TaobaoMobilePromotionCouponApplyResponse struct {
    XMLName xml.Name `xml:"mobile_promotion_coupon_apply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 优惠券领取结果
    CouponApplyResult   *CouponApplyResult `json:"coupon_apply_result,omitempty" xml:"coupon_apply_result,omitempty"`
}
