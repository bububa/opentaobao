package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取垃圾注册防控结果 APIResponse
alibaba.security.jaq.spamregisterprevention.result.fetch

获取垃圾注册防控结果
*/
type AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqSpamregisterpreventionResultFetchResponse `json:"alibaba_security_jaq_spamregisterprevention_result_fetch_response,omitempty"`
}

type AlibabaSecurityJaqSpamregisterpreventionResultFetchResponse struct {

    // 账号风控返回结果
    JaqAccountRiskResult   *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty"`

}
