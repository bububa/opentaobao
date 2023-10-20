package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountResetcodeSendAPIRequest 百川发送找回密码验证码 API请求
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
type TaobaoBaichuanOpenaccountResetcodeSendAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountResetcodeSendRequest 初始化TaobaoBaichuanOpenaccountResetcodeSendAPIRequest对象
func NewTaobaoBaichuanOpenaccountResetcodeSendRequest() *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest {
	return &TaobaoBaichuanOpenaccountResetcodeSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.resetcode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountResetcodeSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountResetcodeSendRequest()
	},
}

// GetTaobaoBaichuanOpenaccountResetcodeSendRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountResetcodeSendAPIRequest
func GetTaobaoBaichuanOpenaccountResetcodeSendAPIRequest() *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest {
	return poolTaobaoBaichuanOpenaccountResetcodeSendAPIRequest.Get().(*TaobaoBaichuanOpenaccountResetcodeSendAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountResetcodeSendAPIRequest 将 TaobaoBaichuanOpenaccountResetcodeSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountResetcodeSendAPIRequest(v *TaobaoBaichuanOpenaccountResetcodeSendAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountResetcodeSendAPIRequest.Put(v)
}
