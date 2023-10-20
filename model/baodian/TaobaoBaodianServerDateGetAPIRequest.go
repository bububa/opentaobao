package baodian

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianServerDateGetAPIRequest 服务器时间获取 API请求
// taobao.baodian.server.date.get
//
// 获取服务器时间
type TaobaoBaodianServerDateGetAPIRequest struct {
	model.Params
}

// NewTaobaoBaodianServerDateGetRequest 初始化TaobaoBaodianServerDateGetAPIRequest对象
func NewTaobaoBaodianServerDateGetRequest() *TaobaoBaodianServerDateGetAPIRequest {
	return &TaobaoBaodianServerDateGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaodianServerDateGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaodianServerDateGetAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.server.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaodianServerDateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaodianServerDateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoBaodianServerDateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaodianServerDateGetRequest()
	},
}

// GetTaobaoBaodianServerDateGetRequest 从 sync.Pool 获取 TaobaoBaodianServerDateGetAPIRequest
func GetTaobaoBaodianServerDateGetAPIRequest() *TaobaoBaodianServerDateGetAPIRequest {
	return poolTaobaoBaodianServerDateGetAPIRequest.Get().(*TaobaoBaodianServerDateGetAPIRequest)
}

// ReleaseTaobaoBaodianServerDateGetAPIRequest 将 TaobaoBaodianServerDateGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaodianServerDateGetAPIRequest(v *TaobaoBaodianServerDateGetAPIRequest) {
	v.Reset()
	poolTaobaoBaodianServerDateGetAPIRequest.Put(v)
}
