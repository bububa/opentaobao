package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户同意用户逆向申请 APIResponse
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请
*/
type AlibabaTclsAelophyRefundAgreeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTclsAelophyRefundAgreeResponse `json:"alibaba_tcls_aelophy_refund_agree_response,omitempty"`
}

type AlibabaTclsAelophyRefundAgreeResponse struct {

    // 出参
    Result   *AlibabaTclsAelophyRefundAgreeApiResult `json:"result,omitempty"`

}
