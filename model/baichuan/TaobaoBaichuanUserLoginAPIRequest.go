package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanuserloginAPIRequest 百川H5登录 API请求
// taobao.baichuan.user.login
//
// 百川H5登录
type TaobaobaichuanuserloginAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanuserloginRequest 初始化TaobaobaichuanuserloginAPIRequest对象
func NewTaobaobaichuanuserloginRequest() *TaobaobaichuanuserloginAPIRequest {
	return &TaobaobaichuanuserloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanuserloginAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanuserloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanuserloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanuserloginAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanuserloginAPIRequest) GetName() string {
	return r._name
}
