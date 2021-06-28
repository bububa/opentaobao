package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取认证会话token APIResponse
alibaba.security.jaq.rp.getverifytoken

聚安全实人认证获取认证会话token
*/
type AlibabaSecurityJaqRpGetverifytokenAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpGetverifytokenResponse
}

type AlibabaSecurityJaqRpGetverifytokenResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_getverifytoken_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // token信息
    
    Data   *RpInitResultBo `json:"data,omitempty" xml:"data,omitempty"`

    
}
