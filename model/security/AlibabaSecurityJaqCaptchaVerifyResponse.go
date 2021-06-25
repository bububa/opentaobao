package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查接口 APIResponse
alibaba.security.jaq.captcha.verify

聚安全安全验证检查
*/
type AlibabaSecurityJaqCaptchaVerifyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqCaptchaVerifyResponse `json:"alibaba_security_jaq_captcha_verify_response,omitempty"`
}

type AlibabaSecurityJaqCaptchaVerifyResponse struct {

    // 出参结构体
    Data   *JaqVerifyCaptchaResult `json:"data,omitempty"`

}
