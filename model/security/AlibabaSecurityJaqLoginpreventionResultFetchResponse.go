package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登录保护结果 APIResponse
alibaba.security.jaq.loginprevention.result.fetch

获取登录保护结果
*/
type AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_loginprevention_result_fetch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 账号风控返回结果
    
    JaqAccountRiskResult   *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"