package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取虚假注册保护结果 APIResponse
alibaba.security.jaq.spamregisterprevention.result.fetch.new

获取虚假注册保护结果
*/
type AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_spamregisterprevention_result_fetch_new_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 账号风控返回结果
    
    JaqAccountRiskResult   *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"