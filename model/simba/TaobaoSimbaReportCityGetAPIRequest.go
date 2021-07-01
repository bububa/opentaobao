package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaReportCityGetAPIRequest
获取城市维度报表 API请求
taobao.simba.report.city.get

获取城市维度报表 */
type TaobaoSimbaReportCityGetAPIRequest struct {
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

// NewTaobaoSimbaReportCityGetRequest 初始化TaobaoSimbaReportCityGetAPIRequest对象
func NewTaobaoSimbaReportCityGetRequest() *TaobaoSimbaReportCityGetAPIRequest {
	return &TaobaoSimbaReportCityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaReportCityGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.report.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaReportCityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 昵称
func (r *TaobaoSimbaReportCityGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is TheDate Setter
// 时间
func (r *TaobaoSimbaReportCityGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// Get TheDate Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// Set is Hour Setter
// 当前小时
func (r *TaobaoSimbaReportCityGetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// Get Hour Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetHour() string {
	return r._hour
}

// Set is CampaignId Setter
// 计划id
func (r *TaobaoSimbaReportCityGetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetCampaignId() string {
	return r._campaignId
}
