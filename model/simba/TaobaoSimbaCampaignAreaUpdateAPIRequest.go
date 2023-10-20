package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignareaupdateAPIRequest 更新一个推广计划的投放地域 API请求
// taobao.simba.campaign.area.update
//
// 更新一个推广计划的投放地域
type TaobaosimbacampaignareaupdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
	_area string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaosimbacampaignareaupdateRequest 初始化TaobaosimbacampaignareaupdateAPIRequest对象
func NewTaobaosimbacampaignareaupdateRequest() *TaobaosimbacampaignareaupdateAPIRequest {
	return &TaobaosimbacampaignareaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignareaupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.area.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignareaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignareaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignareaupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignareaupdateAPIRequest) GetNick() string {
	return r._nick
}

// SetArea is Area Setter
// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
func (r *TaobaosimbacampaignareaupdateAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r TaobaosimbacampaignareaupdateAPIRequest) GetArea() string {
	return r._area
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignareaupdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignareaupdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
