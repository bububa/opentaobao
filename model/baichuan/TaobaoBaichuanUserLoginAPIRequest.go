package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanUserLoginAPIRequest
百川H5登录 API请求
taobao.baichuan.user.login

百川H5登录 */
type TaobaoBaichuanUserLoginAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanUserLoginRequest 初始化TaobaoBaichuanUserLoginAPIRequest对象
func NewTaobaoBaichuanUserLoginRequest() *TaobaoBaichuanUserLoginAPIRequest {
	return &TaobaoBaichuanUserLoginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLoginAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLoginAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// name
func (r *TaobaoBaichuanUserLoginAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoBaichuanUserLoginAPIRequest) GetName() string {
	return r._name
}
