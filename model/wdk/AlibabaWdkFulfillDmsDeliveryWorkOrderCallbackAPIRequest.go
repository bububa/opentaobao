package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest 末端配配送作业回传 API请求
// alibaba.wdk.fulfill.dms.delivery.work.order.callback
//
// 末端配配送作业回传。
type AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DeliveryCallbackOrder
}

// NewAlibabawdkfulfilldmsdeliveryworkordercallbackRequest 初始化AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest对象
func NewAlibabawdkfulfilldmsdeliveryworkordercallbackRequest() *AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest {
	return &AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.dms.delivery.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest) SetCallbackOrder(_callbackOrder *DeliveryCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest) GetCallbackOrder() *DeliveryCallbackOrder {
	return r._callbackOrder
}
