package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountregistercodecheckAPIRequest 百川检查注册验证码 API请求
// taobao.baichuan.openaccount.registercode.check
//
// 百川检查注册验证码
type TaobaobaichuanopenaccountregistercodecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountregistercodecheckRequest 初始化TaobaobaichuanopenaccountregistercodecheckAPIRequest对象
func NewTaobaobaichuanopenaccountregistercodecheckRequest() *TaobaobaichuanopenaccountregistercodecheckAPIRequest {
	return &TaobaobaichuanopenaccountregistercodecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountregistercodecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.registercode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountregistercodecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountregistercodecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountregistercodecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountregistercodecheckAPIRequest) GetName() string {
	return r._name
}
