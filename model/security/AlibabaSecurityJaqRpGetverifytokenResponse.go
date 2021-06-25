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
    Response *AlibabaSecurityJaqRpGetverifytokenResponse `json:"alibaba_security_jaq_rp_getverifytoken_response,omitempty"`
}

type AlibabaSecurityJaqRpGetverifytokenResponse struct {

    // token信息
    Data   *RpInitResultBo `json:"data,omitempty"`

}
