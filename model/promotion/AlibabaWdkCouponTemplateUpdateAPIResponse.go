package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版修改 API返回值 
alibaba.wdk.coupon.template.update

优惠券模版修改
*/
type AlibabaWdkCouponTemplateUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponTemplateUpdateAPIResponseModel
}

// 优惠券模版修改 成功返回结果
type AlibabaWdkCouponTemplateUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_template_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaWdkCouponTemplateUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
