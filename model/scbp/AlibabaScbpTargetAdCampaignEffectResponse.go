package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划维度推广效果 APIResponse
alibaba.scbp.target.ad.campaign.effect

定向推广-获取计划维度推广效果
*/
type AlibabaScbpTargetAdCampaignEffectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdCampaignEffectResponse `json:"alibaba_scbp_target_ad_campaign_effect_response,omitempty"`
}

type AlibabaScbpTargetAdCampaignEffectResponse struct {

    // 数据列表
    EffectList   []TopP4pQuickCampaignEffectView `json:"effect_list,omitempty"`

}
