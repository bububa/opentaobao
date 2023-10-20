package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOrderurlGetAPIRequest 百川订单详情 API请求
// taobao.baichuan.orderurl.get
//
// 百川订单详情
type TaobaoBaichuanOrderurlGetAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOrderurlGetRequest 初始化TaobaoBaichuanOrderurlGetAPIRequest对象
func NewTaobaoBaichuanOrderurlGetRequest() *TaobaoBaichuanOrderurlGetAPIRequest {
	return &TaobaoBaichuanOrderurlGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOrderurlGetAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.orderurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOrderurlGetAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOrderurlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOrderurlGetRequest()
	},
}

// GetTaobaoBaichuanOrderurlGetRequest 从 sync.Pool 获取 TaobaoBaichuanOrderurlGetAPIRequest
func GetTaobaoBaichuanOrderurlGetAPIRequest() *TaobaoBaichuanOrderurlGetAPIRequest {
	return poolTaobaoBaichuanOrderurlGetAPIRequest.Get().(*TaobaoBaichuanOrderurlGetAPIRequest)
}

// ReleaseTaobaoBaichuanOrderurlGetAPIRequest 将 TaobaoBaichuanOrderurlGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOrderurlGetAPIRequest(v *TaobaoBaichuanOrderurlGetAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOrderurlGetAPIRequest.Put(v)
}
