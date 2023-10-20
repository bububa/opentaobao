package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkwarehouseordercancelAPIRequest 仓作业取消下发 API请求
// wdk.warehouse.order.cancel
//
// 仓作业取消下发
type WdkwarehouseordercancelAPIRequest struct {
	model.Params
	// 取消请求
	_cancelRequest *CancelRequest
}

// NewWdkwarehouseordercancelRequest 初始化WdkwarehouseordercancelAPIRequest对象
func NewWdkwarehouseordercancelRequest() *WdkwarehouseordercancelAPIRequest {
	return &WdkwarehouseordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkwarehouseordercancelAPIRequest) GetApiMethodName() string {
	return "wdk.warehouse.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkwarehouseordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkwarehouseordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 取消请求
func (r *WdkwarehouseordercancelAPIRequest) SetCancelRequest(_cancelRequest *CancelRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r WdkwarehouseordercancelAPIRequest) GetCancelRequest() *CancelRequest {
	return r._cancelRequest
}
