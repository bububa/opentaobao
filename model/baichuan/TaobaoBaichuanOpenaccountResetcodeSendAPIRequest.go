package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountresetcodesendAPIRequest 百川发送找回密码验证码 API请求
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
type TaobaobaichuanopenaccountresetcodesendAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountresetcodesendRequest 初始化TaobaobaichuanopenaccountresetcodesendAPIRequest对象
func NewTaobaobaichuanopenaccountresetcodesendRequest() *TaobaobaichuanopenaccountresetcodesendAPIRequest {
	return &TaobaobaichuanopenaccountresetcodesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountresetcodesendAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.resetcode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountresetcodesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountresetcodesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountresetcodesendAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountresetcodesendAPIRequest) GetName() string {
	return r._name
}
