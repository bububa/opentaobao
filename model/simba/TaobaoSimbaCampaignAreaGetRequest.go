package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放地域设置 APIRequest
taobao.simba.campaign.area.get

取得一个推广计划的投放地域设置
*/
type TaobaoSimbaCampaignAreaGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广计划Id
    campaignId   int64 

}

func NewTaobaoSimbaCampaignAreaGetRequest() *TaobaoSimbaCampaignAreaGetRequest{
    return &TaobaoSimbaCampaignAreaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignAreaGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.area.get"
}

func (r TaobaoSimbaCampaignAreaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignAreaGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignAreaGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCampaignAreaGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignAreaGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

