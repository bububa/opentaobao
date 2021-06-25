package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版终止 APIResponse
alibaba.wdk.coupon.template.terminate

优惠券模版终止
*/
type AlibabaWdkCouponTemplateTerminateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponTemplateTerminateResponse `json:"alibaba_wdk_coupon_template_terminate_response,omitempty"`
}

type AlibabaWdkCouponTemplateTerminateResponse struct {

    // 结果
    Result   *AlibabaWdkCouponTemplateTerminateApiResult `json:"result,omitempty"`

}
