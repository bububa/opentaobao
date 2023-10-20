package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest 百川发送注册验证码 API请求
// taobao.baichuan.openaccount.registercode.send
//
// 百川发送注册验证码
type TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountRegistercodeSendRequest 初始化TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest对象
func NewTaobaoBaichuanOpenaccountRegistercodeSendRequest() *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest {
	return &TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.registercode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountRegistercodeSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountRegistercodeSendRequest()
	},
}

// GetTaobaoBaichuanOpenaccountRegistercodeSendRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest
func GetTaobaoBaichuanOpenaccountRegistercodeSendAPIRequest() *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest {
	return poolTaobaoBaichuanOpenaccountRegistercodeSendAPIRequest.Get().(*TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountRegistercodeSendAPIRequest 将 TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountRegistercodeSendAPIRequest(v *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountRegistercodeSendAPIRequest.Put(v)
}
