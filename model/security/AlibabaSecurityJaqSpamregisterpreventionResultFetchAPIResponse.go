package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponse 获取垃圾注册防控结果 API返回值
// alibaba.security.jaq.spamregisterprevention.result.fetch
//
// 获取垃圾注册防控结果
type AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponseModel
}

// AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponseModel is 获取垃圾注册防控结果 成功返回结果
type AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_spamregisterprevention_result_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账号风控返回结果
	JaqAccountRiskResult *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"jaq_account_risk_result,omitempty"`
}
