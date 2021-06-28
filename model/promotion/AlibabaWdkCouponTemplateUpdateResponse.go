package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版修改 APIResponse
alibaba.wdk.coupon.template.update

优惠券模版修改
*/
type AlibabaWdkCouponTemplateUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponTemplateUpdateResponse
}

type AlibabaWdkCouponTemplateUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_template_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponTemplateUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
