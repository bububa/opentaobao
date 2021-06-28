package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查结果获取接口 APIResponse
alibaba.security.jaq.captcha.verify.result.fetch

获取二次验证的结果
*/
type AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqCaptchaVerifyResultFetchResponse
}

type AlibabaSecurityJaqCaptchaVerifyResultFetchResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_captcha_verify_result_fetch_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 验证检查结果结构体
    
    Data   *JaqSecondCheckResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
