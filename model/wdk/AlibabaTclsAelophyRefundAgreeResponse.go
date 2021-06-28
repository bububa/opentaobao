package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户同意用户逆向申请 APIResponse
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请
*/
type AlibabaTclsAelophyRefundAgreeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_aelophy_refund_agree_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *AlibabaTclsAelophyRefundAgreeApiResult `json:"result,omitempty" xml:"