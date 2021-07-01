package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
删除计划 
taobao.feedflow.item.campaign.delete

删除计划
*/
func TaobaoFeedflowItemCampaignDelete(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignDeleteAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCampaignDeleteAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCampaignDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
