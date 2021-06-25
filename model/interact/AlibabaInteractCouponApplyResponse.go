package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取鉴权接口 APIResponse
alibaba.interact.coupon.apply

鉴权接口，为coupon.apply接口鉴权
*/
type AlibabaInteractCouponApplyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractCouponApplyResponse `json:"alibaba_interact_coupon_apply_response,omitempty"`
}

type AlibabaInteractCouponApplyResponse struct {

    // 无用参数，top限制一定要有出参，增加此参数
    Stub   string `json:"stub,omitempty"`

}
