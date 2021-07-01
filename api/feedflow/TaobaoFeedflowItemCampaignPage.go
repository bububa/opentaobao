package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
批量查询计划列表 
taobao.feedflow.item.campaign.page

批量查询计划列表
*/
func TaobaoFeedflowItemCampaignPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignPageAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCampaignPageAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCampaignPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
