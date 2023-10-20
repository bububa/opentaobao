package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountLoginAPIRequest 百川用户名密码登录 API请求
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
type TaobaoBaichuanOpenaccountLoginAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountLoginRequest 初始化TaobaoBaichuanOpenaccountLoginAPIRequest对象
func NewTaobaoBaichuanOpenaccountLoginRequest() *TaobaoBaichuanOpenaccountLoginAPIRequest {
	return &TaobaoBaichuanOpenaccountLoginAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountLoginAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLoginAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountLoginRequest()
	},
}

// GetTaobaoBaichuanOpenaccountLoginRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountLoginAPIRequest
func GetTaobaoBaichuanOpenaccountLoginAPIRequest() *TaobaoBaichuanOpenaccountLoginAPIRequest {
	return poolTaobaoBaichuanOpenaccountLoginAPIRequest.Get().(*TaobaoBaichuanOpenaccountLoginAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountLoginAPIRequest 将 TaobaoBaichuanOpenaccountLoginAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountLoginAPIRequest(v *TaobaoBaichuanOpenaccountLoginAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountLoginAPIRequest.Put(v)
}
