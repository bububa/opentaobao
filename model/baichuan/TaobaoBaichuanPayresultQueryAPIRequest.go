package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanPayresultQueryAPIRequest 百川支付完成回调 API请求
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
type TaobaoBaichuanPayresultQueryAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanPayresultQueryRequest 初始化TaobaoBaichuanPayresultQueryAPIRequest对象
func NewTaobaoBaichuanPayresultQueryRequest() *TaobaoBaichuanPayresultQueryAPIRequest {
	return &TaobaoBaichuanPayresultQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanPayresultQueryAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.payresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanPayresultQueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanPayresultQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanPayresultQueryRequest()
	},
}

// GetTaobaoBaichuanPayresultQueryRequest 从 sync.Pool 获取 TaobaoBaichuanPayresultQueryAPIRequest
func GetTaobaoBaichuanPayresultQueryAPIRequest() *TaobaoBaichuanPayresultQueryAPIRequest {
	return poolTaobaoBaichuanPayresultQueryAPIRequest.Get().(*TaobaoBaichuanPayresultQueryAPIRequest)
}

// ReleaseTaobaoBaichuanPayresultQueryAPIRequest 将 TaobaoBaichuanPayresultQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanPayresultQueryAPIRequest(v *TaobaoBaichuanPayresultQueryAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanPayresultQueryAPIRequest.Put(v)
}
