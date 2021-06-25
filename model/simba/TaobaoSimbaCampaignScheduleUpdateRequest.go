package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的分时折扣设置 APIRequest
taobao.simba.campaign.schedule.update

更新一个推广计划的分时折扣设置
*/
type TaobaoSimbaCampaignScheduleUpdateRequest struct {
    model.Params

    // 推广计划Id
    campaignId   int64 

    // 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比；
    schedule   string 

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaCampaignScheduleUpdateRequest() *TaobaoSimbaCampaignScheduleUpdateRequest{
    return &TaobaoSimbaCampaignScheduleUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.schedule.update"
}

func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignScheduleUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaCampaignScheduleUpdateRequest) SetSchedule(schedule string) error {
    r.schedule = schedule
    r.Set("schedule", schedule)
    return nil
}

func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetSchedule() string {
    return r.schedule
}

func (r *TaobaoSimbaCampaignScheduleUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetNick() string {
    return r.nick
}

