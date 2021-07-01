package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版查询 API返回值 
alibaba.wdk.coupon.template.query

优惠券模版查询
*/
type AlibabaWdkCouponTemplateQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponTemplateQueryAPIResponseModel
}

// 优惠券模版查询 成功返回结果
type AlibabaWdkCouponTemplateQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_template_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaWdkCouponTemplateQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
