package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
修改计划 
alibaba.scbp.ad.campaign.update

修改计划
*/
func AlibabaScbpAdCampaignUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignUpdateRequest, session string) (*scbp.AlibabaScbpAdCampaignUpdateAPIResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
