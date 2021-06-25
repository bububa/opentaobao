package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放平台设置 APIRequest
taobao.simba.campaign.platform.get

获得一个推广计划的投放平台设置
*/
type TaobaoSimbaCampaignPlatformGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广计划Id
    campaignId   int64 

}

func NewTaobaoSimbaCampaignPlatformGetRequest() *TaobaoSimbaCampaignPlatformGetRequest{
    return &TaobaoSimbaCampaignPlatformGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignPlatformGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.platform.get"
}

func (r TaobaoSimbaCampaignPlatformGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignPlatformGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignPlatformGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCampaignPlatformGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignPlatformGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

