package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取推广计划实时报表数据 
taobao.simba.rtrpt.campaign.get

获取推广计划实时报表数据
*/
func TaobaoSimbaRtrptCampaignGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCampaignGetRequest, session string) (*simba.TaobaoSimbaRtrptCampaignGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRtrptCampaignGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
