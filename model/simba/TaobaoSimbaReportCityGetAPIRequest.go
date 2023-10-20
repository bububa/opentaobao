package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbareportcitygetAPIRequest 获取城市维度报表 API请求
// taobao.simba.report.city.get
//
// 获取城市维度报表
type TaobaosimbareportcitygetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
	// 计划id
	_campaignId string
}

// NewTaobaosimbareportcitygetRequest 初始化TaobaosimbareportcitygetAPIRequest对象
func NewTaobaosimbareportcitygetRequest() *TaobaosimbareportcitygetAPIRequest {
	return &TaobaosimbareportcitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbareportcitygetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.report.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbareportcitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbareportcitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbareportcitygetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbareportcitygetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaosimbareportcitygetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbareportcitygetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaosimbareportcitygetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaosimbareportcitygetAPIRequest) GetHour() string {
	return r._hour
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaosimbareportcitygetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbareportcitygetAPIRequest) GetCampaignId() string {
	return r._campaignId
}
