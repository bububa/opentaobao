package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全-实人认证日志打点接口 APIResponse
alibaba.security.jaq.rp.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpRphitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpRphitResponse `json:"alibaba_security_jaq_rp_rphit_response,omitempty"` 
    AlibabaSecurityJaqRpRphitResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpRphitResponse struct {

    // data
    
    Data   string `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpRphitResponse struct {

    // data
    Data   string `json:"data,omitempty"`

}
