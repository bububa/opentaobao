package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeordersuccesscreateAPIRequest 五道口终态订单创建 API请求
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
type AlibabawdktradeordersuccesscreateAPIRequest struct {
	model.Params
	// 创单请求对象
	_orderRequest *OrderSuccessRequest
}

// NewAlibabawdktradeordersuccesscreateRequest 初始化AlibabawdktradeordersuccesscreateAPIRequest对象
func NewAlibabawdktradeordersuccesscreateRequest() *AlibabawdktradeordersuccesscreateAPIRequest {
	return &AlibabawdktradeordersuccesscreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktradeordersuccesscreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.success.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktradeordersuccesscreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktradeordersuccesscreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 创单请求对象
func (r *AlibabawdktradeordersuccesscreateAPIRequest) SetOrderRequest(_orderRequest *OrderSuccessRequest) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r AlibabawdktradeordersuccesscreateAPIRequest) GetOrderRequest() *OrderSuccessRequest {
	return r._orderRequest
}
