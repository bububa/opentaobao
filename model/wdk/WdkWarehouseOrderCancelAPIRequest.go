package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWarehouseOrderCancelAPIRequest) GetApiMethodName() string {
	return "wdk.warehouse.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWarehouseOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
