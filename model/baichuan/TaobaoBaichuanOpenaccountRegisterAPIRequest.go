package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountregisterAPIRequest 百川账号注册 API请求
// taobao.baichuan.openaccount.register
//
// 百川账号注册
type TaobaobaichuanopenaccountregisterAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountregisterRequest 初始化TaobaobaichuanopenaccountregisterAPIRequest对象
func NewTaobaobaichuanopenaccountregisterRequest() *TaobaobaichuanopenaccountregisterAPIRequest {
	return &TaobaobaichuanopenaccountregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountregisterAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountregisterAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountregisterAPIRequest) GetName() string {
	return r._name
}
