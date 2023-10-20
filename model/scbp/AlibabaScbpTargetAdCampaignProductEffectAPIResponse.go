package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdCampaignProductEffectAPIResponse 定向推广-获取计划中产品推广效果 API返回值
// alibaba.scbp.target.ad.campaign.product.effect
//
// 定向推广-获取计划中产品推广效果
type AlibabaScbpTargetAdCampaignProductEffectAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdCampaignProductEffectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel is 定向推广-获取计划中产品推广效果 成功返回结果
type AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_campaign_product_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品数据
	ProductList []TopP4pQuickProductEffectView `json:"product_list,omitempty" xml:"product_list>top_p4p_quick_product_effect_view,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductList = m.ProductList[:0]
	m.TotalPage = 0
	m.TotalNum = 0
}

var poolAlibabaScbpTargetAdCampaignProductEffectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdCampaignProductEffectAPIResponse)
	},
}

// GetAlibabaScbpTargetAdCampaignProductEffectAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdCampaignProductEffectAPIResponse
func GetAlibabaScbpTargetAdCampaignProductEffectAPIResponse() *AlibabaScbpTargetAdCampaignProductEffectAPIResponse {
	return poolAlibabaScbpTargetAdCampaignProductEffectAPIResponse.Get().(*AlibabaScbpTargetAdCampaignProductEffectAPIResponse)
}

// ReleaseAlibabaScbpTargetAdCampaignProductEffectAPIResponse 将 AlibabaScbpTargetAdCampaignProductEffectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdCampaignProductEffectAPIResponse(v *AlibabaScbpTargetAdCampaignProductEffectAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdCampaignProductEffectAPIResponse.Put(v)
}
