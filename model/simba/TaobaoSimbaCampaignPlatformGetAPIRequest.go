package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignplatformgetAPIRequest 取得一个推广计划的投放平台设置 API请求
// taobao.simba.campaign.platform.get
//
// 获得一个推广计划的投放平台设置
type TaobaosimbacampaignplatformgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaosimbacampaignplatformgetRequest 初始化TaobaosimbacampaignplatformgetAPIRequest对象
func NewTaobaosimbacampaignplatformgetRequest() *TaobaosimbacampaignplatformgetAPIRequest {
	return &TaobaosimbacampaignplatformgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignplatformgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.platform.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignplatformgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignplatformgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignplatformgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignplatformgetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignplatformgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignplatformgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
