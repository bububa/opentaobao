package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-获取计划维度推广效果 
alibaba.scbp.target.ad.campaign.effect

定向推广-获取计划维度推广效果
*/
func AlibabaScbpTargetAdCampaignEffect(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdCampaignEffectAPIRequest, session string) (*scbp.AlibabaScbpTargetAdCampaignEffectAPIResponse, error) {
    var resp scbp.AlibabaScbpTargetAdCampaignEffectAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
