package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadcampaigneffectAPIResponse 定向推广-获取计划维度推广效果 API返回值
// alibaba.scbp.target.ad.campaign.effect
//
// 定向推广-获取计划维度推广效果
type AlibabascbptargetadcampaigneffectAPIResponse struct {
	model.CommonResponse
	AlibabascbptargetadcampaigneffectAPIResponseModel
}

// AlibabascbptargetadcampaigneffectAPIResponseModel is 定向推广-获取计划维度推广效果 成功返回结果
type AlibabascbptargetadcampaigneffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_campaign_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据列表
	EffectList []TopP4pQuickCampaignEffectView `json:"effect_list,omitempty" xml:"effect_list>top_p4p_quick_campaign_effect_view,omitempty"`
}
