package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdCampaignEffectAPIResponse 定向推广-获取计划维度推广效果 API返回值
// alibaba.scbp.target.ad.campaign.effect
//
// 定向推广-获取计划维度推广效果
type AlibabaScbpTargetAdCampaignEffectAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdCampaignEffectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdCampaignEffectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdCampaignEffectAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdCampaignEffectAPIResponseModel is 定向推广-获取计划维度推广效果 成功返回结果
type AlibabaScbpTargetAdCampaignEffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_campaign_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据列表
	EffectList []TopP4pQuickCampaignEffectView `json:"effect_list,omitempty" xml:"effect_list>top_p4p_quick_campaign_effect_view,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdCampaignEffectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EffectList = m.EffectList[:0]
}

var poolAlibabaScbpTargetAdCampaignEffectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdCampaignEffectAPIResponse)
	},
}

// GetAlibabaScbpTargetAdCampaignEffectAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdCampaignEffectAPIResponse
func GetAlibabaScbpTargetAdCampaignEffectAPIResponse() *AlibabaScbpTargetAdCampaignEffectAPIResponse {
	return poolAlibabaScbpTargetAdCampaignEffectAPIResponse.Get().(*AlibabaScbpTargetAdCampaignEffectAPIResponse)
}

// ReleaseAlibabaScbpTargetAdCampaignEffectAPIResponse 将 AlibabaScbpTargetAdCampaignEffectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdCampaignEffectAPIResponse(v *AlibabaScbpTargetAdCampaignEffectAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdCampaignEffectAPIResponse.Put(v)
}
