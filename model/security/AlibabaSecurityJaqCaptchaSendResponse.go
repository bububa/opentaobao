package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证发起接口 APIResponse
alibaba.security.jaq.captcha.send

聚安全安全验证发起
*/
type AlibabaSecurityJaqCaptchaSendAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqCaptchaSendResponse
}

type AlibabaSecurityJaqCaptchaSendResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_captcha_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参结构体
    
    Data   *JaqSendCaptchaResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
