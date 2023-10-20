package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaHourReportCampaignGetAPIRequest 计划维度小时报表获取 API请求
// taobao.simba.hour.report.campaign.get
//
// 计划维度小时报表获取
type TaobaoSimbaHourReportCampaignGetAPIRequest struct {
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

// NewTaobaoSimbaHourReportCampaignGetRequest 初始化TaobaoSimbaHourReportCampaignGetAPIRequest对象
func NewTaobaoSimbaHourReportCampaignGetRequest() *TaobaoSimbaHourReportCampaignGetAPIRequest {
	return &TaobaoSimbaHourReportCampaignGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaHourReportCampaignGetAPIRequest) Reset() {
	r._nick = ""
	r._theDate = ""
	r._hour = ""
	r._campaignId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.hour.report.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaHourReportCampaignGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportCampaignGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportCampaignGetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetHour() string {
	return r._hour
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoSimbaHourReportCampaignGetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaHourReportCampaignGetAPIRequest) GetCampaignId() string {
	return r._campaignId
}

var poolTaobaoSimbaHourReportCampaignGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaHourReportCampaignGetRequest()
	},
}

// GetTaobaoSimbaHourReportCampaignGetRequest 从 sync.Pool 获取 TaobaoSimbaHourReportCampaignGetAPIRequest
func GetTaobaoSimbaHourReportCampaignGetAPIRequest() *TaobaoSimbaHourReportCampaignGetAPIRequest {
	return poolTaobaoSimbaHourReportCampaignGetAPIRequest.Get().(*TaobaoSimbaHourReportCampaignGetAPIRequest)
}

// ReleaseTaobaoSimbaHourReportCampaignGetAPIRequest 将 TaobaoSimbaHourReportCampaignGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaHourReportCampaignGetAPIRequest(v *TaobaoSimbaHourReportCampaignGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaHourReportCampaignGetAPIRequest.Put(v)
}
