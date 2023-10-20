package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest 百川检查注册验证码 API请求
// taobao.baichuan.openaccount.registercode.check
//
// 百川检查注册验证码
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountRegistercodeCheckRequest 初始化TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest对象
func NewTaobaoBaichuanOpenaccountRegistercodeCheckRequest() *TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest {
	return &TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.registercode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountRegistercodeCheckRequest()
	},
}

// GetTaobaoBaichuanOpenaccountRegistercodeCheckRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest
func GetTaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest() *TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest {
	return poolTaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest.Get().(*TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest 将 TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest(v *TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest.Put(v)
}
