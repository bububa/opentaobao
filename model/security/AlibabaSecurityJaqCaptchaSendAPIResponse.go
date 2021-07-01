package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqCaptchaSendAPIResponse
聚安全安全验证发起接口 API返回值
alibaba.security.jaq.captcha.send

聚安全安全验证发起 */
type AlibabaSecurityJaqCaptchaSendAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqCaptchaSendAPIResponseModel
}

// AlibabaSecurityJaqCaptchaSendAPIResponseModel is 聚安全安全验证发起接口 成功返回结果
type AlibabaSecurityJaqCaptchaSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_captcha_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqSendCaptchaResult `json:"data,omitempty" xml:"data,omitempty"`
}
