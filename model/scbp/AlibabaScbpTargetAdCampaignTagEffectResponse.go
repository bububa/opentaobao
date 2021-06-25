package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划定向效果数据 APIResponse
alibaba.scbp.target.ad.campaign.tag.effect

定向推广-获取推广计划定向效果数据
*/
type AlibabaScbpTargetAdCampaignTagEffectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdCampaignTagEffectResponse `json:"alibaba_scbp_target_ad_campaign_tag_effect_response,omitempty"`
}

type AlibabaScbpTargetAdCampaignTagEffectResponse struct {

    // 定向标签数据
    EffectList   []TopP4pQuickTagEffectView `json:"effect_list,omitempty"`

    // 总条数
    TotalNum   int64 `json:"total_num,omitempty"`

}
