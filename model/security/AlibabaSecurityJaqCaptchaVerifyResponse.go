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
    // Response *AlibabaSecurityJaqCaptchaVerifyResponse `json:"alibaba_security_jaq_captcha_verify_response,omitempty"` 
    AlibabaSecurityJaqCaptchaVerifyResponse
}

/* model for simplify = false
type AlibabaSecurityJaqCaptchaVerifyResponse struct {

    // 出参结构体
    
    Data  *struct {
        JaqVerifyCaptchaResult  *JaqVerifyCaptchaResult `json:"jaq_verify_captcha_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqCaptchaVerifyResponse struct {

    // 出参结构体
    Data   *JaqVerifyCaptchaResult `json:"data,omitempty"`

}
