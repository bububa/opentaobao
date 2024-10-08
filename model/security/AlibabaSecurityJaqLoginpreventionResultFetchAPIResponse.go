package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse 获取登录保护结果 API返回值
// alibaba.security.jaq.loginprevention.result.fetch
//
// 获取登录保护结果
type AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqLoginpreventionResultFetchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqLoginpreventionResultFetchAPIResponseModel).Reset()
}

// AlibabaSecurityJaqLoginpreventionResultFetchAPIResponseModel is 获取登录保护结果 成功返回结果
type AlibabaSecurityJaqLoginpreventionResultFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_loginprevention_result_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账号风控返回结果
	JaqAccountRiskResult *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"jaq_account_risk_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqLoginpreventionResultFetchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.JaqAccountRiskResult = nil
}

var poolAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse)
	},
}

// GetAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse
func GetAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse() *AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse {
	return poolAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse.Get().(*AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse)
}

// ReleaseAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse 将 AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse(v *AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqLoginpreventionResultFetchAPIResponse.Put(v)
}
