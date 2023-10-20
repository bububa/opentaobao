package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegisterAPIRequest 百川账号注册 API请求
// taobao.baichuan.openaccount.register
//
// 百川账号注册
type TaobaoBaichuanOpenaccountRegisterAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountRegisterRequest 初始化TaobaoBaichuanOpenaccountRegisterAPIRequest对象
func NewTaobaoBaichuanOpenaccountRegisterRequest() *TaobaoBaichuanOpenaccountRegisterAPIRequest {
	return &TaobaoBaichuanOpenaccountRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountRegisterAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegisterAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountRegisterRequest()
	},
}

// GetTaobaoBaichuanOpenaccountRegisterRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountRegisterAPIRequest
func GetTaobaoBaichuanOpenaccountRegisterAPIRequest() *TaobaoBaichuanOpenaccountRegisterAPIRequest {
	return poolTaobaoBaichuanOpenaccountRegisterAPIRequest.Get().(*TaobaoBaichuanOpenaccountRegisterAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountRegisterAPIRequest 将 TaobaoBaichuanOpenaccountRegisterAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountRegisterAPIRequest(v *TaobaoBaichuanOpenaccountRegisterAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountRegisterAPIRequest.Put(v)
}
