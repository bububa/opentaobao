package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
重发核销码 APIResponse
tmall.servicecenter.workcard.verifycode.resend

重发核销码
*/
type TmallServicecenterWorkcardVerifycodeResendAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardVerifycodeResendResponse `json:"tmall_servicecenter_workcard_verifycode_resend_response,omitempty"` 
    TmallServicecenterWorkcardVerifycodeResendResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardVerifycodeResendResponse struct {

    // 调用结果
    
    Result  *struct {
        TmallServicecenterWorkcardVerifycodeResendResult  *TmallServicecenterWorkcardVerifycodeResendResult `json:"tmall_servicecenter_workcard_verifycode_resend_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardVerifycodeResendResponse struct {

    // 调用结果
    Result   *TmallServicecenterWorkcardVerifycodeResendResult `json:"result,omitempty"`

}
