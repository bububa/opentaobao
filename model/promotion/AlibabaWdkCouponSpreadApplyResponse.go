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
    // Response *AlibabaWdkCouponSpreadApplyResponse `json:"alibaba_wdk_coupon_spread_apply_response,omitempty"` 
    AlibabaWdkCouponSpreadApplyResponse
}

/* model for simplify = false
type AlibabaWdkCouponSpreadApplyResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkCouponSpreadApplyApiResult  *AlibabaWdkCouponSpreadApplyApiResult `json:"alibaba_wdk_coupon_spread_apply_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponSpreadApplyResponse struct {

    // 结果
    Result   *AlibabaWdkCouponSpreadApplyApiResult `json:"result,omitempty"`

}
