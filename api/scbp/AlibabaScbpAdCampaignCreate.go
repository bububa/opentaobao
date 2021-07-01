package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
创建计划 
alibaba.scbp.ad.campaign.create

创建计划
*/
func AlibabaScbpAdCampaignCreate(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignCreateAPIRequest, session string) (*scbp.AlibabaScbpAdCampaignCreateAPIResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
