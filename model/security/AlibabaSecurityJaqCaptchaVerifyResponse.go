package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查接口 APIResponse
alibaba.security.jaq.captcha.verify

聚安全安全验证检查
*/
type AlibabaSecurityJaqCaptchaVerifyAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqCaptchaVerifyResponse
}

type AlibabaSecurityJaqCaptchaVerifyResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_captcha_verify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参结构体
    
    Data   *JaqVerifyCaptchaResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
