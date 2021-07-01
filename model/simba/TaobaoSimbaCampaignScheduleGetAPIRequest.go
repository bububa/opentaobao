package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignScheduleGetAPIRequest
取得一个推广计划的分时折扣设置 API请求
taobao.simba.campaign.schedule.get

取得一个推广计划的分时折扣设置 */
type TaobaoSimbaCampaignScheduleGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignScheduleGetRequest 初始化TaobaoSimbaCampaignScheduleGetAPIRequest对象
func NewTaobaoSimbaCampaignScheduleGetRequest() *TaobaoSimbaCampaignScheduleGetAPIRequest {
	return &TaobaoSimbaCampaignScheduleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.schedule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignScheduleGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignScheduleGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
