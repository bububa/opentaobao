package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取虚假注册保护结果 APIResponse
alibaba.security.jaq.spamregisterprevention.result.fetch.new

获取虚假注册保护结果
*/
type AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewResponse `json:"alibaba_security_jaq_spamregisterprevention_result_fetch_new_response,omitempty"`
}

type AlibabaSecurityJaqSpamregisterpreventionResultFetchNewResponse struct {

    // 账号风控返回结果
    JaqAccountRiskResult   *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty"`

}
