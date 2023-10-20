package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillwarehouseworkordercallbackAPIRequest 标准化仓作业单回传接口 API请求
// alibaba.wdk.fulfill.warehouse.work.order.callback
//
// 标准化仓作业单回传接口
type AlibabawdkfulfillwarehouseworkordercallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DrfHalfDayCcCallbackOrder
}

// NewAlibabawdkfulfillwarehouseworkordercallbackRequest 初始化AlibabawdkfulfillwarehouseworkordercallbackAPIRequest对象
func NewAlibabawdkfulfillwarehouseworkordercallbackRequest() *AlibabawdkfulfillwarehouseworkordercallbackAPIRequest {
	return &AlibabawdkfulfillwarehouseworkordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillwarehouseworkordercallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.warehouse.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillwarehouseworkordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillwarehouseworkordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabawdkfulfillwarehouseworkordercallbackAPIRequest) SetCallbackOrder(_callbackOrder *DrfHalfDayCcCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabawdkfulfillwarehouseworkordercallbackAPIRequest) GetCallbackOrder() *DrfHalfDayCcCallbackOrder {
	return r._callbackOrder
}
