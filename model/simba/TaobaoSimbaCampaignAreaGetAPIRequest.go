package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaGetAPIRequest 取得一个推广计划的投放地域设置 API请求
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
type TaobaoSimbaCampaignAreaGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignAreaGetRequest 初始化TaobaoSimbaCampaignAreaGetAPIRequest对象
func NewTaobaoSimbaCampaignAreaGetRequest() *TaobaoSimbaCampaignAreaGetAPIRequest {
	return &TaobaoSimbaCampaignAreaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.area.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAreaGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignAreaGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
