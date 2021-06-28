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
    // Response *AlibabaScbpTargetAdCampaignEffectResponse `json:"alibaba_scbp_target_ad_campaign_effect_response,omitempty"` 
    AlibabaScbpTargetAdCampaignEffectResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdCampaignEffectResponse struct {

    // 数据列表
    
    EffectList  struct {
        TopP4pQuickCampaignEffectView  []TopP4pQuickCampaignEffectView `json:"top_p4p_quick_campaign_effect_view,omitempty"`
    } `json:"effect_list,omitempty"`
    

}
*/

type AlibabaScbpTargetAdCampaignEffectResponse struct {

    // 数据列表
    EffectList   []TopP4pQuickCampaignEffectView `json:"effect_list,omitempty"`

}
