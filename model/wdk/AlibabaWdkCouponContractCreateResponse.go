package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
营销券合同创建接口 APIResponse
alibaba.wdk.coupon.contract.create

营销券合同创建接口
*/
type AlibabaWdkCouponContractCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkCouponContractCreateResponse `json:"alibaba_wdk_coupon_contract_create_response,omitempty"` 
    AlibabaWdkCouponContractCreateResponse
}

/* model for simplify = false
type AlibabaWdkCouponContractCreateResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaWdkCouponContractCreateApiResult  *AlibabaWdkCouponContractCreateApiResult `json:"alibaba_wdk_coupon_contract_create_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponContractCreateResponse struct {

    // 返回结果
    Result   *AlibabaWdkCouponContractCreateApiResult `json:"result,omitempty"`

}
