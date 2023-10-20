package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignScheduleUpdateAPIRequest 更新一个推广计划的分时折扣设置 API请求
// taobao.simba.campaign.schedule.update
//
// 更新一个推广计划的分时折扣设置
type TaobaoSimbaCampaignScheduleUpdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比；
	_schedule string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignScheduleUpdateRequest 初始化TaobaoSimbaCampaignScheduleUpdateAPIRequest对象
func NewTaobaoSimbaCampaignScheduleUpdateRequest() *TaobaoSimbaCampaignScheduleUpdateAPIRequest {
	return &TaobaoSimbaCampaignScheduleUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignScheduleUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.schedule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignScheduleUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignScheduleUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignScheduleUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignScheduleUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetSchedule is Schedule Setter
// 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比；
func (r *TaobaoSimbaCampaignScheduleUpdateAPIRequest) SetSchedule(_schedule string) error {
	r._schedule = _schedule
	r.Set("schedule", _schedule)
	return nil
}

// GetSchedule Schedule Getter
func (r TaobaoSimbaCampaignScheduleUpdateAPIRequest) GetSchedule() string {
	return r._schedule
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignScheduleUpdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignScheduleUpdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
