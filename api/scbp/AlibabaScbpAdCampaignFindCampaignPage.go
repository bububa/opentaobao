package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
分页查询计划 
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划
*/
func AlibabaScbpAdCampaignFindCampaignPage(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindCampaignPageRequest, session string) (*scbp.AlibabaScbpAdCampaignFindCampaignPageAPIResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignFindCampaignPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
