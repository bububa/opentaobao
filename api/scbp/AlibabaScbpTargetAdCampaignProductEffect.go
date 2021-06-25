package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-获取计划中产品推广效果 
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果
*/
func AlibabaScbpTargetAdCampaignProductEffect(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdCampaignProductEffectRequest, session string) (*scbp.AlibabaScbpTargetAdCampaignProductEffectResponse, error) {
    var resp scbp.AlibabaScbpTargetAdCampaignProductEffectAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
