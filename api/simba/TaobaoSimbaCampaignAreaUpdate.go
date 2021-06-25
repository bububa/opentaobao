package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
更新一个推广计划的投放地域 
taobao.simba.campaign.area.update

更新一个推广计划的投放地域
*/
func TaobaoSimbaCampaignAreaUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAreaUpdateRequest, session string) (*simba.TaobaoSimbaCampaignAreaUpdateResponse, error) {
    var resp simba.TaobaoSimbaCampaignAreaUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
