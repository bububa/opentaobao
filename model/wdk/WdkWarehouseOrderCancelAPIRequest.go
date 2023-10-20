package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkWarehouseOrderCancelAPIRequest 仓作业取消下发 API请求
// wdk.warehouse.order.cancel
//
// 仓作业取消下发
type WdkWarehouseOrderCancelAPIRequest struct {
	model.Params
	// 取消请求
	_cancelRequest *CancelRequest
}

// NewWdkWarehouseOrderCancelRequest 初始化WdkWarehouseOrderCancelAPIRequest对象
func NewWdkWarehouseOrderCancelRequest() *WdkWarehouseOrderCancelAPIRequest {
	return &WdkWarehouseOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkWarehouseOrderCancelAPIRequest) Reset() {
	r._cancelRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWarehouseOrderCancelAPIRequest) GetApiMethodName() string {
	return "wdk.warehouse.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWarehouseOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkWarehouseOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 取消请求
func (r *WdkWarehouseOrderCancelAPIRequest) SetCancelRequest(_cancelRequest *CancelRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r WdkWarehouseOrderCancelAPIRequest) GetCancelRequest() *CancelRequest {
	return r._cancelRequest
}

var poolWdkWarehouseOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkWarehouseOrderCancelRequest()
	},
}

// GetWdkWarehouseOrderCancelRequest 从 sync.Pool 获取 WdkWarehouseOrderCancelAPIRequest
func GetWdkWarehouseOrderCancelAPIRequest() *WdkWarehouseOrderCancelAPIRequest {
	return poolWdkWarehouseOrderCancelAPIRequest.Get().(*WdkWarehouseOrderCancelAPIRequest)
}

// ReleaseWdkWarehouseOrderCancelAPIRequest 将 WdkWarehouseOrderCancelAPIRequest 放入 sync.Pool
func ReleaseWdkWarehouseOrderCancelAPIRequest(v *WdkWarehouseOrderCancelAPIRequest) {
	v.Reset()
	poolWdkWarehouseOrderCancelAPIRequest.Put(v)
}
