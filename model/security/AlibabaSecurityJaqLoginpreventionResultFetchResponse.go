package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取登录保护结果 APIResponse
alibaba.security.jaq.loginprevention.result.fetch

获取登录保护结果
*/
type AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqLoginpreventionResultFetchResponse `json:"alibaba_security_jaq_loginprevention_result_fetch_response,omitempty"` 
    AlibabaSecurityJaqLoginpreventionResultFetchResponse
}

/* model for simplify = false
type AlibabaSecurityJaqLoginpreventionResultFetchResponse struct {

    // 账号风控返回结果
    
    JaqAccountRiskResult  *struct {
        JaqAccountRiskResult  *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty"`
    } `json:"jaq_account_risk_result,omitempty"`
    

}
*/

type AlibabaSecurityJaqLoginpreventionResultFetchResponse struct {

    // 账号风控返回结果
    JaqAccountRiskResult   *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty"`

}
