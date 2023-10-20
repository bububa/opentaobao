package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse 获取活动保护结果 API返回值
// alibaba.security.jaq.campaignprevention.result.fetch
//
// 获取活动保护结果
type AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponseModel).Reset()
}

// AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponseModel is 获取活动保护结果 成功返回结果
type AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_campaignprevention_result_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账号风控返回结果
	JaqAccountRiskResult *JaqAccountRiskResult `json:"jaq_account_risk_result,omitempty" xml:"jaq_account_risk_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.JaqAccountRiskResult = nil
}

var poolAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse)
	},
}

// GetAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse
func GetAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse() *AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse {
	return poolAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse.Get().(*AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse)
}

// ReleaseAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse 将 AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse(v *AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse.Put(v)
}
