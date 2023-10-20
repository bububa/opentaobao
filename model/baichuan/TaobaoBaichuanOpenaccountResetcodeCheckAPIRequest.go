package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountresetcodecheckAPIRequest 百川验证找回密码验证码 API请求
// taobao.baichuan.openaccount.resetcode.check
//
// 百川验证找回密码验证码
type TaobaobaichuanopenaccountresetcodecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountresetcodecheckRequest 初始化TaobaobaichuanopenaccountresetcodecheckAPIRequest对象
func NewTaobaobaichuanopenaccountresetcodecheckRequest() *TaobaobaichuanopenaccountresetcodecheckAPIRequest {
	return &TaobaobaichuanopenaccountresetcodecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountresetcodecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.resetcode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountresetcodecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountresetcodecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountresetcodecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountresetcodecheckAPIRequest) GetName() string {
	return r._name
}
