package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
取得一个推广计划的分时折扣设置 
taobao.simba.campaign.schedule.get

取得一个推广计划的分时折扣设置
*/
func TaobaoSimbaCampaignScheduleGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignScheduleGetRequest, session string) (*simba.TaobaoSimbaCampaignScheduleGetResponse, error) {
    var resp simba.TaobaoSimbaCampaignScheduleGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
