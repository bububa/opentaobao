package baichuan

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountPasswordResetAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.password.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountPasswordResetAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountPasswordResetAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountPasswordResetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountPasswordResetRequest()
	},
}

// GetTaobaoBaichuanOpenaccountPasswordResetRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountPasswordResetAPIRequest
func GetTaobaoBaichuanOpenaccountPasswordResetAPIRequest() *TaobaoBaichuanOpenaccountPasswordResetAPIRequest {
	return poolTaobaoBaichuanOpenaccountPasswordResetAPIRequest.Get().(*TaobaoBaichuanOpenaccountPasswordResetAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountPasswordResetAPIRequest 将 TaobaoBaichuanOpenaccountPasswordResetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountPasswordResetAPIRequest(v *TaobaoBaichuanOpenaccountPasswordResetAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountPasswordResetAPIRequest.Put(v)
}
