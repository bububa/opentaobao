package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的日限额 APIRequest
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

func NewTaobaoSimbaCampaignBudgetGetRequest() *TaobaoSimbaCampaignBudgetGetRequest{
    return &TaobaoSimbaCampaignBudgetGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignBudgetGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.budget.get"
}

func (r TaobaoSimbaCampaignBudgetGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignBudgetGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignBudgetGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCampaignBudgetGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignBudgetGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

