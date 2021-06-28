package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
更新一个推广计划的日限额 
taobao.simba.campaign.budget.update

更新一个推广计划的日限额
*/
func TaobaoSimbaCampaignBudgetUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignBudgetUpdateRequest, session string) (*simba.TaobaoSimbaCampaignBudgetUpdateAPIResponse, error) {
    var resp simba.TaobaoSimbaCampaignBudgetUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
