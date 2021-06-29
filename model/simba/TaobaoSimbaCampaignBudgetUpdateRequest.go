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
    _useSmooth   bool
    // 推广计划Id
    _campaignId   int64
    // 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
    _budget   int64
    // 主人昵称
    _nick   string
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
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetUseSmooth(_useSmooth bool) error {
    r._useSmooth = _useSmooth
    r.Set("use_smooth", _useSmooth)
    return nil
}

// UseSmooth Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetUseSmooth() bool {
    return r._useSmooth
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetCampaignId() int64 {
    return r._campaignId
}
// Budget Setter
// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetBudget(_budget int64) error {
    r._budget = _budget
    r.Set("budget", _budget)
    return nil
}

// Budget Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetBudget() int64 {
    return r._budget
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignBudgetUpdateRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignBudgetUpdateRequest) GetNick() string {
    return r._nick
}
