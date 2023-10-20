package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleorderdummysendAPIRequest 闲鱼无需物流发货 API请求
// alibaba.idle.order.dummy.send
//
// 适用于电子卡券等虚拟商品不需要物流的商品发货。
type AlibabaidleorderdummysendAPIRequest struct {
	model.Params
	// 请求
	_paramOrderDummySendRequest *OrderDummySendRequest
}

// NewAlibabaidleorderdummysendRequest 初始化AlibabaidleorderdummysendAPIRequest对象
func NewAlibabaidleorderdummysendRequest() *AlibabaidleorderdummysendAPIRequest {
	return &AlibabaidleorderdummysendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleorderdummysendAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.order.dummy.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleorderdummysendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleorderdummysendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderDummySendRequest is ParamOrderDummySendRequest Setter
// 请求
func (r *AlibabaidleorderdummysendAPIRequest) SetParamOrderDummySendRequest(_paramOrderDummySendRequest *OrderDummySendRequest) error {
	r._paramOrderDummySendRequest = _paramOrderDummySendRequest
	r.Set("param_order_dummy_send_request", _paramOrderDummySendRequest)
	return nil
}

// GetParamOrderDummySendRequest ParamOrderDummySendRequest Getter
func (r AlibabaidleorderdummysendAPIRequest) GetParamOrderDummySendRequest() *OrderDummySendRequest {
	return r._paramOrderDummySendRequest
}
