package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
销量明星跟新预算相关接口 
taobao.simba.salestar.campaign.budget.update

更新一个推广计划的日限额
*/
func TaobaoSimbaSalestarCampaignBudgetUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCampaignBudgetUpdateRequest, session string) (*simba.TaobaoSimbaSalestarCampaignBudgetUpdateResponse, error) {
    var resp simba.TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
