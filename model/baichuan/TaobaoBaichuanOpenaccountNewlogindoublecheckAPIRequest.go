package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest 百川新登录二次验证 API请求
// taobao.baichuan.openaccount.newlogindoublecheck
//
// 百川新登录二次验证
type TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountnewlogindoublecheckRequest 初始化TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest对象
func NewTaobaobaichuanopenaccountnewlogindoublecheckRequest() *TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest {
	return &TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.newlogindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest) GetName() string {
	return r._name
}
