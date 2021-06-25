package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星跟新预算相关接口 APIRequest
taobao.simba.salestar.campaign.budget.update

更新一个推广计划的日限额
*/
type TaobaoSimbaSalestarCampaignBudgetUpdateRequest struct {
    model.Params

    // 推广计划Id
    campaignId   int64 

    // 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
    budget   int64 

}

func NewTaobaoSimbaSalestarCampaignBudgetUpdateRequest() *TaobaoSimbaSalestarCampaignBudgetUpdateRequest{
    return &TaobaoSimbaSalestarCampaignBudgetUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.campaign.budget.update"
}

func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarCampaignBudgetUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaSalestarCampaignBudgetUpdateRequest) SetBudget(budget int64) error {
    r.budget = budget
    r.Set("budget", budget)
    return nil
}

func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetBudget() int64 {
    return r.budget
}

