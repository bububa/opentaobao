package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
更新一个推广计划的分时折扣设置 
taobao.simba.campaign.schedule.update

更新一个推广计划的分时折扣设置
*/
func TaobaoSimbaCampaignScheduleUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignScheduleUpdateRequest, session string) (*simba.TaobaoSimbaCampaignScheduleUpdateResponse, error) {
    var resp simba.TaobaoSimbaCampaignScheduleUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
