package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkRefundAggregateAPIRequest 淘鲜达退款单按门店聚合查询 API请求
// alibaba.wdk.refund.aggregate
//
// 淘鲜达退款单按门店聚合查询
type AlibabaWdkRefundAggregateAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundAggregateQueryRequest *RefundAggregateQueryRequest
}

// NewAlibabaWdkRefundAggregateRequest 初始化AlibabaWdkRefundAggregateAPIRequest对象
func NewAlibabaWdkRefundAggregateRequest() *AlibabaWdkRefundAggregateAPIRequest {
	return &AlibabaWdkRefundAggregateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkRefundAggregateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.refund.aggregate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkRefundAggregateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkRefundAggregateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundAggregateQueryRequest is RefundAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkRefundAggregateAPIRequest) SetRefundAggregateQueryRequest(_refundAggregateQueryRequest *RefundAggregateQueryRequest) error {
	r._refundAggregateQueryRequest = _refundAggregateQueryRequest
	r.Set("refund_aggregate_query_request", _refundAggregateQueryRequest)
	return nil
}

// GetRefundAggregateQueryRequest RefundAggregateQueryRequest Getter
func (r AlibabaWdkRefundAggregateAPIRequest) GetRefundAggregateQueryRequest() *RefundAggregateQueryRequest {
	return r._refundAggregateQueryRequest
}
