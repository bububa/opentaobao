package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse 保护结果反馈 API返回值
// alibaba.security.jaq.spamregisterprevention.feedback
//
// 用户通过这个接口对垃圾注册防控结果进行反馈
type AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponseModel).Reset()
}

// AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponseModel is 保护结果反馈 成功返回结果
type AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_spamregisterprevention_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// feedBack返回结果
	FeedBackResult *JaqFeedBackResult `json:"feed_back_result,omitempty" xml:"feed_back_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FeedBackResult = nil
}

var poolAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse)
	},
}

// GetAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse
func GetAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse() *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse {
	return poolAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse.Get().(*AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse)
}

// ReleaseAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse 将 AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse(v *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse.Put(v)
}
