package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划定向效果数据 API返回值 
alibaba.scbp.target.ad.campaign.tag.effect

定向推广-获取推广计划定向效果数据
*/
type AlibabaScbpTargetAdCampaignTagEffectAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdCampaignTagEffectAPIResponseModel
}

// 定向推广-获取推广计划定向效果数据 成功返回结果
type AlibabaScbpTargetAdCampaignTagEffectAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_campaign_tag_effect_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 定向标签数据
    EffectList   []TopP4pQuickTagEffectView `json:"effect_list,omitempty" xml:"effect_list>top_p4p_quick_tag_effect_view,omitempty"`
    // 总条数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
