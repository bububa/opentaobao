package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbartrptcustgetAPIRequest 获取账户实时报表数据 API请求
// taobao.simba.rtrpt.cust.get
//
// 获取账户实时报表数据
type TaobaosimbartrptcustgetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// NewTaobaosimbartrptcustgetRequest 初始化TaobaosimbartrptcustgetAPIRequest对象
func NewTaobaosimbartrptcustgetRequest() *TaobaosimbartrptcustgetAPIRequest {
	return &TaobaosimbartrptcustgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbartrptcustgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.cust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbartrptcustgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbartrptcustgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbartrptcustgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbartrptcustgetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaosimbartrptcustgetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbartrptcustgetAPIRequest) GetTheDate() string {
	return r._theDate
}
