package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的分时折扣设置 API请求
taobao.simba.campaign.schedule.get

取得一个推广计划的分时折扣设置
*/
type TaobaoSimbaCampaignScheduleGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广计划Id
    campaignId   int64
}

// 初始化TaobaoSimbaCampaignScheduleGetRequest对象
func NewTaobaoSimbaCampaignScheduleGetRequest() *TaobaoSimbaCampaignScheduleGetRequest{
    return &TaobaoSimbaCampaignScheduleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignScheduleGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.schedule.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignScheduleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignScheduleGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignScheduleGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignScheduleGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignScheduleGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
