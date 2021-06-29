package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的日限额 API请求
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

// 初始化TaobaoSimbaCampaignBudgetUpdateRequest对象
func NewTaobaoSimbaCampaignBudgetUpdateRequest() *TaobaoSimbaCampaignBudgetUpdateRequest{
    return &TaobaoSimbaCampaignBudgetUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.budget.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UseSmooth Setter
// 是否平滑消耗：false-否，true-是
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetUseSmooth(useSmooth bool) error {
    r.useSmooth = useSmooth
    r.Set("use_smooth", useSmooth)
    return nil
}

// UseSmooth Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetUseSmooth() bool {
    return r.useSmooth
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}
// Budget Setter
// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetBudget(budget int64) error {
    r.budget = budget
    r.Set("budget", budget)
    return nil
}

// Budget Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetBudget() int64 {
    return r.budget
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetNick() string {
    return r.nick
}
