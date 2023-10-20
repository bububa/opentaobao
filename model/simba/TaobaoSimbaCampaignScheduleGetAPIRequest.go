package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignschedulegetAPIRequest 取得一个推广计划的分时折扣设置 API请求
// taobao.simba.campaign.schedule.get
//
// 取得一个推广计划的分时折扣设置
type TaobaosimbacampaignschedulegetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaosimbacampaignschedulegetRequest 初始化TaobaosimbacampaignschedulegetAPIRequest对象
func NewTaobaosimbacampaignschedulegetRequest() *TaobaosimbacampaignschedulegetAPIRequest {
	return &TaobaosimbacampaignschedulegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignschedulegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.schedule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignschedulegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignschedulegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignschedulegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignschedulegetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignschedulegetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignschedulegetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
