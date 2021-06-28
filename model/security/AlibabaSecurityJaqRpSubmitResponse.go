package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证提交认证接口 APIResponse
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpSubmitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpSubmitResponse `json:"alibaba_security_jaq_rp_submit_response,omitempty"` 
    AlibabaSecurityJaqRpSubmitResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpSubmitResponse struct {

    // 结果信息
    
    Data  *struct {
        RpSubmitResult  *RpSubmitResult `json:"rp_submit_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpSubmitResponse struct {

    // 结果信息
    Data   *RpSubmitResult `json:"data,omitempty"`

}
