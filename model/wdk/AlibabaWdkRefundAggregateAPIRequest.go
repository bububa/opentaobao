package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkrefundaggregateAPIRequest 淘鲜达退款单按门店聚合查询 API请求
// alibaba.wdk.refund.aggregate
//
// 淘鲜达退款单按门店聚合查询
type AlibabawdkrefundaggregateAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundAggregateQueryRequest *RefundAggregateQueryRequest
}

// NewAlibabawdkrefundaggregateRequest 初始化AlibabawdkrefundaggregateAPIRequest对象
func NewAlibabawdkrefundaggregateRequest() *AlibabawdkrefundaggregateAPIRequest {
	return &AlibabawdkrefundaggregateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkrefundaggregateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.refund.aggregate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkrefundaggregateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkrefundaggregateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundAggregateQueryRequest is RefundAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabawdkrefundaggregateAPIRequest) SetRefundAggregateQueryRequest(_refundAggregateQueryRequest *RefundAggregateQueryRequest) error {
	r._refundAggregateQueryRequest = _refundAggregateQueryRequest
	r.Set("refund_aggregate_query_request", _refundAggregateQueryRequest)
	return nil
}

// GetRefundAggregateQueryRequest RefundAggregateQueryRequest Getter
func (r AlibabawdkrefundaggregateAPIRequest) GetRefundAggregateQueryRequest() *RefundAggregateQueryRequest {
	return r._refundAggregateQueryRequest
}
