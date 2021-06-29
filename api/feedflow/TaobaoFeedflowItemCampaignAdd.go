package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流增加推广计划 
taobao.feedflow.item.campaign.add

信息流增加推广计划
*/
func TaobaoFeedflowItemCampaignAdd(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignAddRequest, session string) (*feedflow.TaobaoFeedflowItemCampaignAddAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCampaignAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
