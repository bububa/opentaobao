package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyordergetAPIRequest 翱象拉取订单接口 API请求
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
type AlibabaaelophyordergetAPIRequest struct {
	model.Params
	// 请求对象
	_orderGetRequest *OrderGetRequest
}

// NewAlibabaaelophyordergetRequest 初始化AlibabaaelophyordergetAPIRequest对象
func NewAlibabaaelophyordergetRequest() *AlibabaaelophyordergetAPIRequest {
	return &AlibabaaelophyordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaelophyordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaelophyordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaelophyordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderGetRequest is OrderGetRequest Setter
// 请求对象
func (r *AlibabaaelophyordergetAPIRequest) SetOrderGetRequest(_orderGetRequest *OrderGetRequest) error {
	r._orderGetRequest = _orderGetRequest
	r.Set("order_get_request", _orderGetRequest)
	return nil
}

// GetOrderGetRequest OrderGetRequest Getter
func (r AlibabaaelophyordergetAPIRequest) GetOrderGetRequest() *OrderGetRequest {
	return r._orderGetRequest
}
