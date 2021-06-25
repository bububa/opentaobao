package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的投放地域 APIRequest
taobao.simba.campaign.area.update

更新一个推广计划的投放地域
*/
type TaobaoSimbaCampaignAreaUpdateRequest struct {
    model.Params

    // 推广计划Id
    campaignId   int64 

    // 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
    area   string 

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaCampaignAreaUpdateRequest() *TaobaoSimbaCampaignAreaUpdateRequest{
    return &TaobaoSimbaCampaignAreaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignAreaUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.area.update"
}

func (r TaobaoSimbaCampaignAreaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignAreaUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignAreaUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaCampaignAreaUpdateRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

func (r TaobaoSimbaCampaignAreaUpdateRequest) GetArea() string {
    return r.area
}

func (r *TaobaoSimbaCampaignAreaUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignAreaUpdateRequest) GetNick() string {
    return r.nick
}

