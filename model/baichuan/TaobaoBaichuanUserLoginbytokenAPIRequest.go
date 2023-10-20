package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanUserLoginbytokenAPIRequest 百川手淘信任登录 API请求
// taobao.baichuan.user.loginbytoken
//
// 百川手淘信任登录
type TaobaoBaichuanUserLoginbytokenAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanUserLoginbytokenRequest 初始化TaobaoBaichuanUserLoginbytokenAPIRequest对象
func NewTaobaoBaichuanUserLoginbytokenRequest() *TaobaoBaichuanUserLoginbytokenAPIRequest {
	return &TaobaoBaichuanUserLoginbytokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLoginbytokenAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.loginbytoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLoginbytokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanUserLoginbytokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanUserLoginbytokenAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanUserLoginbytokenAPIRequest) GetName() string {
	return r._name
}
