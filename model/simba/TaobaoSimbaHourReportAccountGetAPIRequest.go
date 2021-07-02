package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaHourReportAccountGetAPIRequest 账户级别小时报表获取 API请求
// taobao.simba.hour.report.account.get
//
// 获取账户小时实时报表数据
type TaobaoSimbaHourReportAccountGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
}

// NewTaobaoSimbaHourReportAccountGetRequest 初始化TaobaoSimbaHourReportAccountGetAPIRequest对象
func NewTaobaoSimbaHourReportAccountGetRequest() *TaobaoSimbaHourReportAccountGetAPIRequest {
	return &TaobaoSimbaHourReportAccountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaHourReportAccountGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.hour.report.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaHourReportAccountGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 昵称
func (r *TaobaoSimbaHourReportAccountGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaHourReportAccountGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportAccountGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// Get TheDate Getter
func (r TaobaoSimbaHourReportAccountGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// Set is Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportAccountGetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// Get Hour Getter
func (r TaobaoSimbaHourReportAccountGetAPIRequest) GetHour() string {
	return r._hour
}
