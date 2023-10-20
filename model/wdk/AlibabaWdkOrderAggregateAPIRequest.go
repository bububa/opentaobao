package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkorderaggregateAPIRequest 淘鲜达订单按门店机台号聚合查询 API请求
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
type AlibabawdkorderaggregateAPIRequest struct {
	model.Params
	// 系统自动生成
	_orderAggregateQueryRequest *OrderAggregateQueryRequest
}

// NewAlibabawdkorderaggregateRequest 初始化AlibabawdkorderaggregateAPIRequest对象
func NewAlibabawdkorderaggregateRequest() *AlibabawdkorderaggregateAPIRequest {
	return &AlibabawdkorderaggregateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkorderaggregateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.aggregate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkorderaggregateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkorderaggregateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderAggregateQueryRequest is OrderAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabawdkorderaggregateAPIRequest) SetOrderAggregateQueryRequest(_orderAggregateQueryRequest *OrderAggregateQueryRequest) error {
	r._orderAggregateQueryRequest = _orderAggregateQueryRequest
	r.Set("order_aggregate_query_request", _orderAggregateQueryRequest)
	return nil
}

// GetOrderAggregateQueryRequest OrderAggregateQueryRequest Getter
func (r AlibabawdkorderaggregateAPIRequest) GetOrderAggregateQueryRequest() *OrderAggregateQueryRequest {
	return r._orderAggregateQueryRequest
}
