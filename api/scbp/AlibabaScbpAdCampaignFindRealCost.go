package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
批量查询计划消耗数据 
alibaba.scbp.ad.campaign.find.real.cost

批量查询计划消耗数据
*/
func AlibabaScbpAdCampaignFindRealCost(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindRealCostRequest, session string) (*scbp.AlibabaScbpAdCampaignFindRealCostResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignFindRealCostAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
