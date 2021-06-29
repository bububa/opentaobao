package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星跟新预算相关接口 API请求
taobao.simba.salestar.campaign.budget.update

更新一个推广计划的日限额
*/
type TaobaoSimbaSalestarCampaignBudgetUpdateRequest struct {
    model.Params
    // 推广计划Id
    _campaignId   int64
    // 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
    _budget   int64
}

// 初始化TaobaoSimbaSalestarCampaignBudgetUpdateRequest对象
func NewTaobaoSimbaSalestarCampaignBudgetUpdateRequest() *TaobaoSimbaSalestarCampaignBudgetUpdateRequest{
    return &TaobaoSimbaSalestarCampaignBudgetUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.campaign.budget.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarCampaignBudgetUpdateRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetCampaignId() int64 {
    return r._campaignId
}
// Budget Setter
// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
func (r *TaobaoSimbaSalestarCampaignBudgetUpdateRequest) SetBudget(_budget int64) error {
    r._budget = _budget
    r.Set("budget", _budget)
    return nil
}

// Budget Getter
func (r TaobaoSimbaSalestarCampaignBudgetUpdateRequest) GetBudget() int64 {
    return r._budget
}
