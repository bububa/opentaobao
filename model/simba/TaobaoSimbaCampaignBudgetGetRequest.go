package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的日限额 API请求
taobao.simba.campaign.budget.get

取得一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetGetAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广计划Id
    _campaignId   int64
}

// 初始化TaobaoSimbaCampaignBudgetGetAPIRequest对象
func NewTaobaoSimbaCampaignBudgetGetRequest() *TaobaoSimbaCampaignBudgetGetAPIRequest{
    return &TaobaoSimbaCampaignBudgetGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.budget.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignBudgetGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignBudgetGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
