package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkorderrefundlistAPIRequest 五道口交易退款批量查询 API请求
// alibaba.wdk.order.refund.list
//
// 按照条件查询退款数据
type AlibabawdkorderrefundlistAPIRequest struct {
	model.Params
	// 查询条件
	_batchQueryRefundRequest *BatchQueryRefundRequest
}

// NewAlibabawdkorderrefundlistRequest 初始化AlibabawdkorderrefundlistAPIRequest对象
func NewAlibabawdkorderrefundlistRequest() *AlibabawdkorderrefundlistAPIRequest {
	return &AlibabawdkorderrefundlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkorderrefundlistAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkorderrefundlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkorderrefundlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchQueryRefundRequest is BatchQueryRefundRequest Setter
// 查询条件
func (r *AlibabawdkorderrefundlistAPIRequest) SetBatchQueryRefundRequest(_batchQueryRefundRequest *BatchQueryRefundRequest) error {
	r._batchQueryRefundRequest = _batchQueryRefundRequest
	r.Set("batch_query_refund_request", _batchQueryRefundRequest)
	return nil
}

// GetBatchQueryRefundRequest BatchQueryRefundRequest Getter
func (r AlibabawdkorderrefundlistAPIRequest) GetBatchQueryRefundRequest() *BatchQueryRefundRequest {
	return r._batchQueryRefundRequest
}
