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
type TaobaoSimbaCampaignBudgetGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广计划Id
    campaignId   int64
}

// 初始化TaobaoSimbaCampaignBudgetGetRequest对象
func NewTaobaoSimbaCampaignBudgetGetRequest() *TaobaoSimbaCampaignBudgetGetRequest{
    return &TaobaoSimbaCampaignBudgetGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignBudgetGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.budget.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignBudgetGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignBudgetGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignBudgetGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignBudgetGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignBudgetGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
