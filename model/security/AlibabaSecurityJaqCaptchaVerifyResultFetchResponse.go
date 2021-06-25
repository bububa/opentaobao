package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查结果获取接口 APIResponse
alibaba.security.jaq.captcha.verify.result.fetch

获取二次验证的结果
*/
type AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqCaptchaVerifyResultFetchResponse `json:"alibaba_security_jaq_captcha_verify_result_fetch_response,omitempty"`
}

type AlibabaSecurityJaqCaptchaVerifyResultFetchResponse struct {

    // 验证检查结果结构体
    Data   *JaqSecondCheckResult `json:"data,omitempty"`

}
