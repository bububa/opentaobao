package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取活动保护结果 APIResponse
alibaba.security.jaq.campaignprevention.result.fetch

获取活动保护结果
*/
type AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqCampaignpreventionResultFetchResponse `json:"alibaba_security_jaq_campaignprevention_result_fetch_response,omitempty"`
}

type AlibabaSecurityJaqCampaignpreventionResultFetchResponse struct {

    // 账号风控返回结果
    JaqAccountRiskResult   *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty"`

}
