package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanUserLoginAPIRequest 百川H5登录 API请求
// taobao.baichuan.user.login
//
// 百川H5登录
type TaobaoBaichuanUserLoginAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanUserLoginRequest 初始化TaobaoBaichuanUserLoginAPIRequest对象
func NewTaobaoBaichuanUserLoginRequest() *TaobaoBaichuanUserLoginAPIRequest {
	return &TaobaoBaichuanUserLoginAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanUserLoginAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLoginAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanUserLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanUserLoginAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanUserLoginAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanUserLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanUserLoginRequest()
	},
}

// GetTaobaoBaichuanUserLoginRequest 从 sync.Pool 获取 TaobaoBaichuanUserLoginAPIRequest
func GetTaobaoBaichuanUserLoginAPIRequest() *TaobaoBaichuanUserLoginAPIRequest {
	return poolTaobaoBaichuanUserLoginAPIRequest.Get().(*TaobaoBaichuanUserLoginAPIRequest)
}

// ReleaseTaobaoBaichuanUserLoginAPIRequest 将 TaobaoBaichuanUserLoginAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanUserLoginAPIRequest(v *TaobaoBaichuanUserLoginAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanUserLoginAPIRequest.Put(v)
}
