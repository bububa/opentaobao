package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户拒绝用户逆向申请 APIResponse
alibaba.tcls.aelophy.refund.disagree

saas 售后逆向 商户拒绝用户逆向申请
*/
type AlibabaTclsAelophyRefundDisagreeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyRefundDisagreeResponse `json:"alibaba_tcls_aelophy_refund_disagree_response,omitempty"` 
    AlibabaTclsAelophyRefundDisagreeResponse
}

/* model for simplify = false
type AlibabaTclsAelophyRefundDisagreeResponse struct {

    // 出参
    
    Result  *struct {
        AlibabaTclsAelophyRefundDisagreeApiResult  *AlibabaTclsAelophyRefundDisagreeApiResult `json:"alibaba_tcls_aelophy_refund_disagree_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaTclsAelophyRefundDisagreeResponse struct {

    // 出参
    Result   *AlibabaTclsAelophyRefundDisagreeApiResult `json:"result,omitempty"`

}
