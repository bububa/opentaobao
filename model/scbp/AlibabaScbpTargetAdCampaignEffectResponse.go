package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划维度推广效果 APIResponse
alibaba.scbp.target.ad.campaign.effect

定向推广-获取计划维度推广效果
*/
type AlibabaScbpTargetAdCampaignEffectAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_campaign_effect_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 数据列表
    
    EffectList   []TopP4pQuickCampaignEffectView `json:"effect_list,omitempty" xml:"