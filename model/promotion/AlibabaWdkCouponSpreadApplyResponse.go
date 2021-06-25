package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
普通发券 APIResponse
alibaba.wdk.coupon.spread.apply

优惠券发放
*/
type AlibabaWdkCouponSpreadApplyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponSpreadApplyResponse `json:"alibaba_wdk_coupon_spread_apply_response,omitempty"`
}

type AlibabaWdkCouponSpreadApplyResponse struct {

    // 结果
    Result   *AlibabaWdkCouponSpreadApplyApiResult `json:"result,omitempty"`

}
