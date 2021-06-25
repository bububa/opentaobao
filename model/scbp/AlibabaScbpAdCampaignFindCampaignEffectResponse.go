package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划效果数据 APIResponse
alibaba.scbp.ad.campaign.find.campaign.effect

批量查询计划效果数据
*/
type AlibabaScbpAdCampaignFindCampaignEffectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdCampaignFindCampaignEffectResponse `json:"alibaba_scbp_ad_campaign_find_campaign_effect_response,omitempty"`
}

type AlibabaScbpAdCampaignFindCampaignEffectResponse struct {

    // json数据，key是计划id(campaignId), value包含开始时间(statDate),曝光(impr),点击(click),消耗(cost),推广时长(onlineMin)
    Result   string `json:"result,omitempty"`

}
