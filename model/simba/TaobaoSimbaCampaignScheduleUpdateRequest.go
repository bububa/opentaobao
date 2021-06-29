package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的分时折扣设置 API请求
taobao.simba.campaign.schedule.update

更新一个推广计划的分时折扣设置
*/
type TaobaoSimbaCampaignScheduleUpdateRequest struct {
    model.Params
    // 推广计划Id
    _campaignId   int64
    // 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比；
    _schedule   string
    // 主人昵称
    _nick   string
}

// 初始化TaobaoSimbaCampaignScheduleUpdateRequest对象
func NewTaobaoSimbaCampaignScheduleUpdateRequest() *TaobaoSimbaCampaignScheduleUpdateRequest{
    return &TaobaoSimbaCampaignScheduleUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.schedule.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignScheduleUpdateRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetCampaignId() int64 {
    return r._campaignId
}
// Schedule Setter
// 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比；
func (r *TaobaoSimbaCampaignScheduleUpdateRequest) SetSchedule(_schedule string) error {
    r._schedule = _schedule
    r.Set("schedule", _schedule)
    return nil
}

// Schedule Getter
func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetSchedule() string {
    return r._schedule
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignScheduleUpdateRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignScheduleUpdateRequest) GetNick() string {
    return r._nick
}
