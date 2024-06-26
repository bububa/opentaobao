package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaReportCityGetAPIRequest 获取城市维度报表 API请求
// taobao.simba.report.city.get
//
// 获取城市维度报表
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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaReportCityGetAPIRequest) Reset() {
	r._nick = ""
	r._theDate = ""
	r._hour = ""
	r._campaignId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaReportCityGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.report.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaReportCityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaReportCityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaReportCityGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaoSimbaReportCityGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaoSimbaReportCityGetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetHour() string {
	return r._hour
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoSimbaReportCityGetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaReportCityGetAPIRequest) GetCampaignId() string {
	return r._campaignId
}

var poolTaobaoSimbaReportCityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaReportCityGetRequest()
	},
}

// GetTaobaoSimbaReportCityGetRequest 从 sync.Pool 获取 TaobaoSimbaReportCityGetAPIRequest
func GetTaobaoSimbaReportCityGetAPIRequest() *TaobaoSimbaReportCityGetAPIRequest {
	return poolTaobaoSimbaReportCityGetAPIRequest.Get().(*TaobaoSimbaReportCityGetAPIRequest)
}

// ReleaseTaobaoSimbaReportCityGetAPIRequest 将 TaobaoSimbaReportCityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaReportCityGetAPIRequest(v *TaobaoSimbaReportCityGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaReportCityGetAPIRequest.Put(v)
}
