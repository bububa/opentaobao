package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaaccountbalancegetAPIRequest 获取实时余额，”元”为单位 API请求
// taobao.simba.account.balance.get
//
// 获取实时余额，”元”为单位
type TaobaosimbaaccountbalancegetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
}

// NewTaobaosimbaaccountbalancegetRequest 初始化TaobaosimbaaccountbalancegetAPIRequest对象
func NewTaobaosimbaaccountbalancegetRequest() *TaobaosimbaaccountbalancegetAPIRequest {
	return &TaobaosimbaaccountbalancegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaaccountbalancegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.account.balance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaaccountbalancegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaaccountbalancegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbaaccountbalancegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaaccountbalancegetAPIRequest) GetNick() string {
	return r._nick
}
