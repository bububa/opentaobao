package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版创建 API返回值 
alibaba.wdk.coupon.template.create

开放给外部商家创建优惠券模版
*/
type AlibabaWdkCouponTemplateCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponTemplateCreateResponse
}

// 优惠券模版创建 成功返回结果
type AlibabaWdkCouponTemplateCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_template_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaWdkCouponTemplateCreateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
