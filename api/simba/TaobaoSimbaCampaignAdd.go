package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
创建一个推广计划 
taobao.simba.campaign.add

创建一个推广计划
*/
func TaobaoSimbaCampaignAdd(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAddRequest, session string) (*simba.TaobaoSimbaCampaignAddResponse, error) {
    var resp simba.TaobaoSimbaCampaignAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
