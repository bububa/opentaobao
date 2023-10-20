package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse 获取虚假注册保护结果 API返回值
// alibaba.security.jaq.spamregisterprevention.result.fetch.new
//
// 获取虚假注册保护结果
type AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponseModel).Reset()
}

// AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponseModel is 获取虚假注册保护结果 成功返回结果
type AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_spamregisterprevention_result_fetch_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账号风控返回结果
	JaqAccountRiskResult *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"jaq_account_risk_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.JaqAccountRiskResult = nil
}

var poolAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse)
	},
}

// GetAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse
func GetAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse() *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse {
	return poolAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse.Get().(*AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse)
}

// ReleaseAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse 将 AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse(v *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse.Put(v)
}
