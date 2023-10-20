package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbahourreportcampaigngetAPIRequest 计划维度小时报表获取 API请求
// taobao.simba.hour.report.campaign.get
//
// 计划维度小时报表获取
type TaobaosimbahourreportcampaigngetAPIRequest struct {
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

// NewTaobaosimbahourreportcampaigngetRequest 初始化TaobaosimbahourreportcampaigngetAPIRequest对象
func NewTaobaosimbahourreportcampaigngetRequest() *TaobaosimbahourreportcampaigngetAPIRequest {
	return &TaobaosimbahourreportcampaigngetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.hour.report.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbahourreportcampaigngetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaosimbahourreportcampaigngetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaosimbahourreportcampaigngetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetHour() string {
	return r._hour
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaosimbahourreportcampaigngetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbahourreportcampaigngetAPIRequest) GetCampaignId() string {
	return r._campaignId
}
