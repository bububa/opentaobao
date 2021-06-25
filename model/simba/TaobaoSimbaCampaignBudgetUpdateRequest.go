package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的日限额 APIRequest
taobao.simba.campaign.budget.update

更新一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetUpdateRequest struct {
    model.Params

    // 是否平滑消耗：false-否，true-是
    useSmooth   bool 

    // 推广计划Id
    campaignId   int64 

    // 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
    budget   int64 

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaCampaignBudgetUpdateRequest() *TaobaoSimbaCampaignBudgetUpdateRequest{
    return &TaobaoSimbaCampaignBudgetUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.budget.update"
}

func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetUseSmooth(useSmooth bool) error {
    r.useSmooth = useSmooth
    r.Set("use_smooth", useSmooth)
    return nil
}

func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetUseSmooth() bool {
    return r.useSmooth
}

func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetBudget(budget int64) error {
    r.budget = budget
    r.Set("budget", budget)
    return nil
}

func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetBudget() int64 {
    return r.budget
}

func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetNick() string {
    return r.nick
}

