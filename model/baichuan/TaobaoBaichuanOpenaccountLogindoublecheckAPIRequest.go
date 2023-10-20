package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountlogindoublecheckAPIRequest 百川登录二次验证 API请求
// taobao.baichuan.openaccount.logindoublecheck
//
// 百川登录二次验证
type TaobaobaichuanopenaccountlogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountlogindoublecheckRequest 初始化TaobaobaichuanopenaccountlogindoublecheckAPIRequest对象
func NewTaobaobaichuanopenaccountlogindoublecheckRequest() *TaobaobaichuanopenaccountlogindoublecheckAPIRequest {
	return &TaobaobaichuanopenaccountlogindoublecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountlogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.logindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountlogindoublecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountlogindoublecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountlogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountlogindoublecheckAPIRequest) GetName() string {
	return r._name
}
