package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanUserLogindoublecheckAPIRequest 百川H5登录二次验证 API请求
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
type TaobaoBaichuanUserLogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanUserLogindoublecheckRequest 初始化TaobaoBaichuanUserLogindoublecheckAPIRequest对象
func NewTaobaoBaichuanUserLogindoublecheckRequest() *TaobaoBaichuanUserLogindoublecheckAPIRequest {
	return &TaobaoBaichuanUserLogindoublecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.logindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanUserLogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetName() string {
	return r._name
}
