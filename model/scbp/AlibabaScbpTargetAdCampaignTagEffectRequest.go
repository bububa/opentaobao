package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划定向效果数据 APIRequest
alibaba.scbp.target.ad.campaign.tag.effect

定向推广-获取推广计划定向效果数据
*/
type AlibabaScbpTargetAdCampaignTagEffectRequest struct {
    model.Params

    // 效果数据
    topP4pQuickEffectQuery   *TopP4pQuickEffectQuery 

}

func NewAlibabaScbpTargetAdCampaignTagEffectRequest() *AlibabaScbpTargetAdCampaignTagEffectRequest{
    return &AlibabaScbpTargetAdCampaignTagEffectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdCampaignTagEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.campaign.tag.effect"
}

func (r AlibabaScbpTargetAdCampaignTagEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdCampaignTagEffectRequest) SetTopP4pQuickEffectQuery(topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
    r.topP4pQuickEffectQuery = topP4pQuickEffectQuery
    r.Set("top_p4p_quick_effect_query", topP4pQuickEffectQuery)
    return nil
}

func (r AlibabaScbpTargetAdCampaignTagEffectRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
    return r.topP4pQuickEffectQuery
}

