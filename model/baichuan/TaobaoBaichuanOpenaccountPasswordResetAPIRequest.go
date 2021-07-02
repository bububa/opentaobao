package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountPasswordResetAPIRequest 百川找回密码 API请求
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
type TaobaoBaichuanOpenaccountPasswordResetAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountPasswordResetRequest 初始化TaobaoBaichuanOpenaccountPasswordResetAPIRequest对象
func NewTaobaoBaichuanOpenaccountPasswordResetRequest() *TaobaoBaichuanOpenaccountPasswordResetAPIRequest {
	return &TaobaoBaichuanOpenaccountPasswordResetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.password.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountPasswordResetAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetName() string {
	return r._name
}
