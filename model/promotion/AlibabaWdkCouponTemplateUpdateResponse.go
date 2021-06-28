package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版修改 APIResponse
alibaba.wdk.coupon.template.update

优惠券模版修改
*/
type AlibabaWdkCouponTemplateUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkCouponTemplateUpdateResponse `json:"alibaba_wdk_coupon_template_update_response,omitempty"` 
    AlibabaWdkCouponTemplateUpdateResponse
}

/* model for simplify = false
type AlibabaWdkCouponTemplateUpdateResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkCouponTemplateUpdateApiResult  *AlibabaWdkCouponTemplateUpdateApiResult `json:"alibaba_wdk_coupon_template_update_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponTemplateUpdateResponse struct {

    // 结果
    Result   *AlibabaWdkCouponTemplateUpdateApiResult `json:"result,omitempty"`

}
