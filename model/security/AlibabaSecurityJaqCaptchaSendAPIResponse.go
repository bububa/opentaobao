package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqcaptchasendAPIResponse 聚安全安全验证发起接口 API返回值
// alibaba.security.jaq.captcha.send
//
// 聚安全安全验证发起
type AlibabasecurityjaqcaptchasendAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqcaptchasendAPIResponseModel
}

// AlibabasecurityjaqcaptchasendAPIResponseModel is 聚安全安全验证发起接口 成功返回结果
type AlibabasecurityjaqcaptchasendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_captcha_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqSendCaptchaResult `json:"data,omitempty" xml:"data,omitempty"`
}
