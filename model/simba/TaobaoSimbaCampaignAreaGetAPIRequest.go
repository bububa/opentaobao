package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignareagetAPIRequest 取得一个推广计划的投放地域设置 API请求
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
type TaobaosimbacampaignareagetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaosimbacampaignareagetRequest 初始化TaobaosimbacampaignareagetAPIRequest对象
func NewTaobaosimbacampaignareagetRequest() *TaobaosimbacampaignareagetAPIRequest {
	return &TaobaosimbacampaignareagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignareagetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.area.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignareagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignareagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignareagetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignareagetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignareagetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignareagetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
