package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
通过计划id查询计划 
taobao.feedflow.item.campaign.get

通过计划id查询计划
*/
func TaobaoFeedflowItemCampaignGet(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignGetRequest, session string) (*feedflow.TaobaoFeedflowItemCampaignGetAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCampaignGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
