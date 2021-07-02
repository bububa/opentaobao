package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignPlatformGetAPIRequest 取得一个推广计划的投放平台设置 API请求
// taobao.simba.campaign.platform.get
//
// 获得一个推广计划的投放平台设置
type TaobaoSimbaCampaignPlatformGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignPlatformGetRequest 初始化TaobaoSimbaCampaignPlatformGetAPIRequest对象
func NewTaobaoSimbaCampaignPlatformGetRequest() *TaobaoSimbaCampaignPlatformGetAPIRequest {
	return &TaobaoSimbaCampaignPlatformGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.platform.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignPlatformGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignPlatformGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
