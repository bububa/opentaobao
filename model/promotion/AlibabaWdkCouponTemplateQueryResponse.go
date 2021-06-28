package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版查询 APIResponse
alibaba.wdk.coupon.template.query

优惠券模版查询
*/
type AlibabaWdkCouponTemplateQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_coupon_template_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponTemplateQueryApiResult `json:"result,omitempty" xml:"