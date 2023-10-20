package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountregistercodesendAPIRequest 百川发送注册验证码 API请求
// taobao.baichuan.openaccount.registercode.send
//
// 百川发送注册验证码
type TaobaobaichuanopenaccountregistercodesendAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountregistercodesendRequest 初始化TaobaobaichuanopenaccountregistercodesendAPIRequest对象
func NewTaobaobaichuanopenaccountregistercodesendRequest() *TaobaobaichuanopenaccountregistercodesendAPIRequest {
	return &TaobaobaichuanopenaccountregistercodesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountregistercodesendAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.registercode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountregistercodesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountregistercodesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountregistercodesendAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountregistercodesendAPIRequest) GetName() string {
	return r._name
}
