package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版查询 APIResponse
alibaba.wdk.coupon.template.query

优惠券模版查询
*/
type AlibabaWdkCouponTemplateQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponTemplateQueryResponse `json:"alibaba_wdk_coupon_template_query_response,omitempty"`
}

type AlibabaWdkCouponTemplateQueryResponse struct {

    // 结果
    Result   *AlibabaWdkCouponTemplateQueryApiResult `json:"result,omitempty"`

}
