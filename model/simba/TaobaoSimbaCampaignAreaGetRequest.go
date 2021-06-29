package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放地域设置 API请求
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

// 初始化TaobaoSimbaCampaignAreaGetRequest对象
func NewTaobaoSimbaCampaignAreaGetRequest() *TaobaoSimbaCampaignAreaGetRequest{
    return &TaobaoSimbaCampaignAreaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.area.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAreaGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignAreaGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignAreaGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignAreaGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
