package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountpasswordresetAPIRequest 百川找回密码 API请求
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
type TaobaobaichuanopenaccountpasswordresetAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountpasswordresetRequest 初始化TaobaobaichuanopenaccountpasswordresetAPIRequest对象
func NewTaobaobaichuanopenaccountpasswordresetRequest() *TaobaobaichuanopenaccountpasswordresetAPIRequest {
	return &TaobaobaichuanopenaccountpasswordresetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountpasswordresetAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.password.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountpasswordresetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountpasswordresetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountpasswordresetAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountpasswordresetAPIRequest) GetName() string {
	return r._name
}
