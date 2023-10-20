package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaHourReportAdgroupGetAPIRequest 推广单元小时级别实时报表查询 API请求
// taobao.simba.hour.report.adgroup.get
//
// 推广单元小时级别实时报表查询
type TaobaoSimbaHourReportAdgroupGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
	// 计划id
	_campaignId string
	// 推广单元id
	_adgroupId string
}

// NewTaobaoSimbaHourReportAdgroupGetRequest 初始化TaobaoSimbaHourReportAdgroupGetAPIRequest对象
func NewTaobaoSimbaHourReportAdgroupGetRequest() *TaobaoSimbaHourReportAdgroupGetAPIRequest {
	return &TaobaoSimbaHourReportAdgroupGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaHourReportAdgroupGetAPIRequest) Reset() {
	r._nick = ""
	r._theDate = ""
	r._hour = ""
	r._campaignId = ""
	r._adgroupId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.hour.report.adgroup.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaHourReportAdgroupGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportAdgroupGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportAdgroupGetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetHour() string {
	return r._hour
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoSimbaHourReportAdgroupGetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetCampaignId() string {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaHourReportAdgroupGetAPIRequest) SetAdgroupId(_adgroupId string) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaHourReportAdgroupGetAPIRequest) GetAdgroupId() string {
	return r._adgroupId
}

var poolTaobaoSimbaHourReportAdgroupGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaHourReportAdgroupGetRequest()
	},
}

// GetTaobaoSimbaHourReportAdgroupGetRequest 从 sync.Pool 获取 TaobaoSimbaHourReportAdgroupGetAPIRequest
func GetTaobaoSimbaHourReportAdgroupGetAPIRequest() *TaobaoSimbaHourReportAdgroupGetAPIRequest {
	return poolTaobaoSimbaHourReportAdgroupGetAPIRequest.Get().(*TaobaoSimbaHourReportAdgroupGetAPIRequest)
}

// ReleaseTaobaoSimbaHourReportAdgroupGetAPIRequest 将 TaobaoSimbaHourReportAdgroupGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaHourReportAdgroupGetAPIRequest(v *TaobaoSimbaHourReportAdgroupGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaHourReportAdgroupGetAPIRequest.Put(v)
}
