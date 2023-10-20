package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOldposOrderCreateAPIRequest 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API请求
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
type AlibabaWdkOldposOrderCreateAPIRequest struct {
	model.Params
	// 入参
	_posOrderCreateRequest *PosOrderCreateRequest
}

// NewAlibabaWdkOldposOrderCreateRequest 初始化AlibabaWdkOldposOrderCreateAPIRequest对象
func NewAlibabaWdkOldposOrderCreateRequest() *AlibabaWdkOldposOrderCreateAPIRequest {
	return &AlibabaWdkOldposOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOldposOrderCreateAPIRequest) Reset() {
	r._posOrderCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.oldpos.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosOrderCreateRequest is PosOrderCreateRequest Setter
// 入参
func (r *AlibabaWdkOldposOrderCreateAPIRequest) SetPosOrderCreateRequest(_posOrderCreateRequest *PosOrderCreateRequest) error {
	r._posOrderCreateRequest = _posOrderCreateRequest
	r.Set("pos_order_create_request", _posOrderCreateRequest)
	return nil
}

// GetPosOrderCreateRequest PosOrderCreateRequest Getter
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetPosOrderCreateRequest() *PosOrderCreateRequest {
	return r._posOrderCreateRequest
}

var poolAlibabaWdkOldposOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOldposOrderCreateRequest()
	},
}

// GetAlibabaWdkOldposOrderCreateRequest 从 sync.Pool 获取 AlibabaWdkOldposOrderCreateAPIRequest
func GetAlibabaWdkOldposOrderCreateAPIRequest() *AlibabaWdkOldposOrderCreateAPIRequest {
	return poolAlibabaWdkOldposOrderCreateAPIRequest.Get().(*AlibabaWdkOldposOrderCreateAPIRequest)
}

// ReleaseAlibabaWdkOldposOrderCreateAPIRequest 将 AlibabaWdkOldposOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOldposOrderCreateAPIRequest(v *AlibabaWdkOldposOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkOldposOrderCreateAPIRequest.Put(v)
}
