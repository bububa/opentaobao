package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanuserlogindoublecheckAPIRequest 百川H5登录二次验证 API请求
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
type TaobaobaichuanuserlogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanuserlogindoublecheckRequest 初始化TaobaobaichuanuserlogindoublecheckAPIRequest对象
func NewTaobaobaichuanuserlogindoublecheckRequest() *TaobaobaichuanuserlogindoublecheckAPIRequest {
	return &TaobaobaichuanuserlogindoublecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanuserlogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.logindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanuserlogindoublecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanuserlogindoublecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanuserlogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanuserlogindoublecheckAPIRequest) GetName() string {
	return r._name
}
