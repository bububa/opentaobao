package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest 百川验证找回密码验证码 API请求
// taobao.baichuan.openaccount.resetcode.check
//
// 百川验证找回密码验证码
type TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountResetcodeCheckRequest 初始化TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest对象
func NewTaobaoBaichuanOpenaccountResetcodeCheckRequest() *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest {
	return &TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.resetcode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountResetcodeCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountResetcodeCheckRequest()
	},
}

// GetTaobaoBaichuanOpenaccountResetcodeCheckRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest
func GetTaobaoBaichuanOpenaccountResetcodeCheckAPIRequest() *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest {
	return poolTaobaoBaichuanOpenaccountResetcodeCheckAPIRequest.Get().(*TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountResetcodeCheckAPIRequest 将 TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountResetcodeCheckAPIRequest(v *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountResetcodeCheckAPIRequest.Put(v)
}
