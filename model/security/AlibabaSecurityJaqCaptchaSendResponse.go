package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证发起接口 APIResponse
alibaba.security.jaq.captcha.send

聚安全安全验证发起
*/
type AlibabaSecurityJaqCaptchaSendAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqCaptchaSendResponse `json:"alibaba_security_jaq_captcha_send_response,omitempty"`
}

type AlibabaSecurityJaqCaptchaSendResponse struct {

    // 出参结构体
    Data   *JaqSendCaptchaResult `json:"data,omitempty"`

}
