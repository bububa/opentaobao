package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后逆向申请渲染获取 APIResponse
alibaba.tcls.aelophy.refund.csapplyrender

提供商家代客售后逆向申请渲染获取的接口
*/
type AlibabaTclsAelophyRefundCsapplyrenderAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyRefundCsapplyrenderResponse `json:"alibaba_tcls_aelophy_refund_csapplyrender_response,omitempty"` 
    AlibabaTclsAelophyRefundCsapplyrenderResponse
}

/* model for simplify = false
type AlibabaTclsAelophyRefundCsapplyrenderResponse struct {

    // 响应结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyRefundCsapplyrenderApiResult  *AlibabaTclsAelophyRefundCsapplyrenderApiResult `json:"alibaba_tcls_aelophy_refund_csapplyrender_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyRefundCsapplyrenderResponse struct {

    // 响应结果
    ApiResult   *AlibabaTclsAelophyRefundCsapplyrenderApiResult `json:"api_result,omitempty"`

}
