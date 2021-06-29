package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划效果数据 APIResponse
alibaba.scbp.ad.campaign.find.campaign.effect

批量查询计划效果数据
*/
type AlibabaScbpAdCampaignFindCampaignEffectAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignFindCampaignEffectResponse
}

type AlibabaScbpAdCampaignFindCampaignEffectResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_campaign_effect_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // json数据，key是计划id(campaignId), value包含开始时间(statDate),曝光(impr),点击(click),消耗(cost),推广时长(onlineMin)
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
