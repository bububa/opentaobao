package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbahourreportaccountgetAPIRequest 账户级别小时报表获取 API请求
// taobao.simba.hour.report.account.get
//
// 获取账户小时实时报表数据
type TaobaosimbahourreportaccountgetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
}

// NewTaobaosimbahourreportaccountgetRequest 初始化TaobaosimbahourreportaccountgetAPIRequest对象
func NewTaobaosimbahourreportaccountgetRequest() *TaobaosimbahourreportaccountgetAPIRequest {
	return &TaobaosimbahourreportaccountgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbahourreportaccountgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.hour.report.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbahourreportaccountgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbahourreportaccountgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbahourreportaccountgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbahourreportaccountgetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaosimbahourreportaccountgetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbahourreportaccountgetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaosimbahourreportaccountgetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaosimbahourreportaccountgetAPIRequest) GetHour() string {
	return r._hour
}
