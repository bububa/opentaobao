package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindCampaignEffectAPIResponse 批量查询计划效果数据 API返回值
// alibaba.scbp.ad.campaign.find.campaign.effect
//
// 批量查询计划效果数据
type AlibabaScbpAdCampaignFindCampaignEffectAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignFindCampaignEffectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindCampaignEffectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignFindCampaignEffectAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignFindCampaignEffectAPIResponseModel is 批量查询计划效果数据 成功返回结果
type AlibabaScbpAdCampaignFindCampaignEffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_campaign_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json数据，key是计划id(campaignId), value包含开始时间(statDate),曝光(impr),点击(click),消耗(cost),推广时长(onlineMin)
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindCampaignEffectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaScbpAdCampaignFindCampaignEffectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignFindCampaignEffectAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignFindCampaignEffectAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignFindCampaignEffectAPIResponse
func GetAlibabaScbpAdCampaignFindCampaignEffectAPIResponse() *AlibabaScbpAdCampaignFindCampaignEffectAPIResponse {
	return poolAlibabaScbpAdCampaignFindCampaignEffectAPIResponse.Get().(*AlibabaScbpAdCampaignFindCampaignEffectAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignFindCampaignEffectAPIResponse 将 AlibabaScbpAdCampaignFindCampaignEffectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindCampaignEffectAPIResponse(v *AlibabaScbpAdCampaignFindCampaignEffectAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindCampaignEffectAPIResponse.Put(v)
}
