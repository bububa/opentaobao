package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountResetcodeSendAPIRequest
百川发送找回密码验证码 API请求
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码 */
type TaobaoBaichuanOpenaccountResetcodeSendAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountResetcodeSendRequest 初始化TaobaoBaichuanOpenaccountResetcodeSendAPIRequest对象
func NewTaobaoBaichuanOpenaccountResetcodeSendRequest() *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest {
	return &TaobaoBaichuanOpenaccountResetcodeSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.resetcode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetName() string {
	return r._name
}
