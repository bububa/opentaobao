package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划定向效果数据 APIResponse
alibaba.scbp.target.ad.campaign.tag.effect

定向推广-获取推广计划定向效果数据
*/
type AlibabaScbpTargetAdCampaignTagEffectAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_campaign_tag_effect_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 定向标签数据
    
    EffectList   []TopP4pQuickTagEffectView `json:"effect_list,omitempty" xml:"