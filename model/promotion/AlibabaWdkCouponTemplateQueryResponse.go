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
    // Response *AlibabaWdkCouponTemplateQueryResponse `json:"alibaba_wdk_coupon_template_query_response,omitempty"` 
    AlibabaWdkCouponTemplateQueryResponse
}

/* model for simplify = false
type AlibabaWdkCouponTemplateQueryResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkCouponTemplateQueryApiResult  *AlibabaWdkCouponTemplateQueryApiResult `json:"alibaba_wdk_coupon_template_query_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponTemplateQueryResponse struct {

    // 结果
    Result   *AlibabaWdkCouponTemplateQueryApiResult `json:"result,omitempty"`

}
