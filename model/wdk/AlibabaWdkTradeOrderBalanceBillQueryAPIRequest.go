package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeorderbalancebillqueryAPIRequest 分页拉取订单数据 API请求
// alibaba.wdk.trade.order.balance.bill.query
//
// 提供接口供外部调用，分页拉取订单数据
type AlibabawdktradeorderbalancebillqueryAPIRequest struct {
	model.Params
	// 入参
	_orderBalanceBillRequest *OrderBalanceBillRequest
}

// NewAlibabawdktradeorderbalancebillqueryRequest 初始化AlibabawdktradeorderbalancebillqueryAPIRequest对象
func NewAlibabawdktradeorderbalancebillqueryRequest() *AlibabawdktradeorderbalancebillqueryAPIRequest {
	return &AlibabawdktradeorderbalancebillqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktradeorderbalancebillqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.balance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktradeorderbalancebillqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktradeorderbalancebillqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderBalanceBillRequest is OrderBalanceBillRequest Setter
// 入参
func (r *AlibabawdktradeorderbalancebillqueryAPIRequest) SetOrderBalanceBillRequest(_orderBalanceBillRequest *OrderBalanceBillRequest) error {
	r._orderBalanceBillRequest = _orderBalanceBillRequest
	r.Set("order_balance_bill_request", _orderBalanceBillRequest)
	return nil
}

// GetOrderBalanceBillRequest OrderBalanceBillRequest Getter
func (r AlibabawdktradeorderbalancebillqueryAPIRequest) GetOrderBalanceBillRequest() *OrderBalanceBillRequest {
	return r._orderBalanceBillRequest
}
