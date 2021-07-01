package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqCaptchaVerifyAPIResponse
聚安全安全验证检查接口 API返回值
alibaba.security.jaq.captcha.verify

聚安全安全验证检查 */
type AlibabaSecurityJaqCaptchaVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqCaptchaVerifyAPIResponseModel
}

// AlibabaSecurityJaqCaptchaVerifyAPIResponseModel is 聚安全安全验证检查接口 成功返回结果
type AlibabaSecurityJaqCaptchaVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_captcha_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqVerifyCaptchaResult `json:"data,omitempty" xml:"data,omitempty"`
}
