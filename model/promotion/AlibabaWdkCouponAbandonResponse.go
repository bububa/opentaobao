package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
废券 APIResponse
alibaba.wdk.coupon.abandon

优惠券废弃
*/
type AlibabaWdkCouponAbandonAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponAbandonResponse `json:"alibaba_wdk_coupon_abandon_response,omitempty"`
}

type AlibabaWdkCouponAbandonResponse struct {

    // 结果
    Result   *AlibabaWdkCouponAbandonApiResult `json:"result,omitempty"`

}
