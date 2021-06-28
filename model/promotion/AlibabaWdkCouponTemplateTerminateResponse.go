package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版终止 APIResponse
alibaba.wdk.coupon.template.terminate

优惠券模版终止
*/
type AlibabaWdkCouponTemplateTerminateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponTemplateTerminateResponse
}

type AlibabaWdkCouponTemplateTerminateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_template_terminate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponTemplateTerminateApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
