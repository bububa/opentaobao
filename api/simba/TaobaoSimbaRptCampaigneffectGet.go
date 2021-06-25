package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广计划效果报表数据对象 
taobao.simba.rpt.campaigneffect.get

推广计划效果报表数据对象
*/
func TaobaoSimbaRptCampaigneffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCampaigneffectGetRequest, session string) (*simba.TaobaoSimbaRptCampaigneffectGetResponse, error) {
    var resp simba.TaobaoSimbaRptCampaigneffectGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
