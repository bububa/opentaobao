package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderAggregateAPIRequest 淘鲜达订单按门店机台号聚合查询 API请求
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
type AlibabaWdkOrderAggregateAPIRequest struct {
	model.Params
	// 系统自动生成
	_orderAggregateQueryRequest *OrderAggregateQueryRequest
}

// NewAlibabaWdkOrderAggregateRequest 初始化AlibabaWdkOrderAggregateAPIRequest对象
func NewAlibabaWdkOrderAggregateRequest() *AlibabaWdkOrderAggregateAPIRequest {
	return &AlibabaWdkOrderAggregateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderAggregateAPIRequest) Reset() {
	r._orderAggregateQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderAggregateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.aggregate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderAggregateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderAggregateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderAggregateQueryRequest is OrderAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkOrderAggregateAPIRequest) SetOrderAggregateQueryRequest(_orderAggregateQueryRequest *OrderAggregateQueryRequest) error {
	r._orderAggregateQueryRequest = _orderAggregateQueryRequest
	r.Set("order_aggregate_query_request", _orderAggregateQueryRequest)
	return nil
}

// GetOrderAggregateQueryRequest OrderAggregateQueryRequest Getter
func (r AlibabaWdkOrderAggregateAPIRequest) GetOrderAggregateQueryRequest() *OrderAggregateQueryRequest {
	return r._orderAggregateQueryRequest
}

var poolAlibabaWdkOrderAggregateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderAggregateRequest()
	},
}

// GetAlibabaWdkOrderAggregateRequest 从 sync.Pool 获取 AlibabaWdkOrderAggregateAPIRequest
func GetAlibabaWdkOrderAggregateAPIRequest() *AlibabaWdkOrderAggregateAPIRequest {
	return poolAlibabaWdkOrderAggregateAPIRequest.Get().(*AlibabaWdkOrderAggregateAPIRequest)
}

// ReleaseAlibabaWdkOrderAggregateAPIRequest 将 AlibabaWdkOrderAggregateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderAggregateAPIRequest(v *AlibabaWdkOrderAggregateAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderAggregateAPIRequest.Put(v)
}
