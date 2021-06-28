package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取认证会话token APIResponse
alibaba.security.jaq.rp.getverifytoken

聚安全实人认证获取认证会话token
*/
type AlibabaSecurityJaqRpGetverifytokenAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpGetverifytokenResponse `json:"alibaba_security_jaq_rp_getverifytoken_response,omitempty"` 
    AlibabaSecurityJaqRpGetverifytokenResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpGetverifytokenResponse struct {

    // token信息
    
    Data  *struct {
        RpInitResultBo  *RpInitResultBo `json:"rp_init_result_bo,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpGetverifytokenResponse struct {

    // token信息
    Data   *RpInitResultBo `json:"data,omitempty"`

}
