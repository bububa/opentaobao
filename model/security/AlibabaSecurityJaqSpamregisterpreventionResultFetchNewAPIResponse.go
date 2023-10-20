package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponse 获取虚假注册保护结果 API返回值
// alibaba.security.jaq.spamregisterprevention.result.fetch.new
//
// 获取虚假注册保护结果
type AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponseModel
}

// AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponseModel is 获取虚假注册保护结果 成功返回结果
type AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_spamregisterprevention_result_fetch_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账号风控返回结果
	JaqAccountRiskResult *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"jaq_account_risk_result,omitempty"`
}
