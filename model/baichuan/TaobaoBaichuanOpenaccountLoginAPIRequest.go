package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountloginAPIRequest 百川用户名密码登录 API请求
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
type TaobaobaichuanopenaccountloginAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountloginRequest 初始化TaobaobaichuanopenaccountloginAPIRequest对象
func NewTaobaobaichuanopenaccountloginRequest() *TaobaobaichuanopenaccountloginAPIRequest {
	return &TaobaobaichuanopenaccountloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountloginAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountloginAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountloginAPIRequest) GetName() string {
	return r._name
}
