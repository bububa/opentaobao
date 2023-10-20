package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse 聚安全安全验证检查结果获取接口 API返回值
// alibaba.security.jaq.captcha.verify.result.fetch
//
// 获取二次验证的结果
type AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponseModel).Reset()
}

// AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponseModel is 聚安全安全验证检查结果获取接口 成功返回结果
type AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_captcha_verify_result_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 验证检查结果结构体
	Data *JaqSecondCheckResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse)
	},
}

// GetAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse
func GetAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse() *AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse {
	return poolAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse.Get().(*AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse)
}

// ReleaseAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse 将 AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse(v *AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse.Put(v)
}
