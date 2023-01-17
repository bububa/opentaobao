package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderSuccessCreateAPIRequest 五道口终态订单创建 API请求
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
type AlibabaWdkTradeOrderSuccessCreateAPIRequest struct {
	model.Params
	// 创单请求对象
	_orderRequest *OrderSuccessRequest
}

// NewAlibabaWdkTradeOrderSuccessCreateRequest 初始化AlibabaWdkTradeOrderSuccessCreateAPIRequest对象
func NewAlibabaWdkTradeOrderSuccessCreateRequest() *AlibabaWdkTradeOrderSuccessCreateAPIRequest {
	return &AlibabaWdkTradeOrderSuccessCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderSuccessCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.success.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderSuccessCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeOrderSuccessCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 创单请求对象
func (r *AlibabaWdkTradeOrderSuccessCreateAPIRequest) SetOrderRequest(_orderRequest *OrderSuccessRequest) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r AlibabaWdkTradeOrderSuccessCreateAPIRequest) GetOrderRequest() *OrderSuccessRequest {
	return r._orderRequest
}
