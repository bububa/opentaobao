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
    // Response *AlibabaWdkCouponAbandonResponse `json:"alibaba_wdk_coupon_abandon_response,omitempty"` 
    AlibabaWdkCouponAbandonResponse
}

/* model for simplify = false
type AlibabaWdkCouponAbandonResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkCouponAbandonApiResult  *AlibabaWdkCouponAbandonApiResult `json:"alibaba_wdk_coupon_abandon_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponAbandonResponse struct {

    // 结果
    Result   *AlibabaWdkCouponAbandonApiResult `json:"result,omitempty"`

}
