package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderRefundListAPIRequest 五道口交易退款批量查询 API请求
// alibaba.wdk.order.refund.list
//
// 按照条件查询退款数据
type AlibabaWdkOrderRefundListAPIRequest struct {
	model.Params
	// 查询条件
	_batchQueryRefundRequest *BatchQueryRefundRequest
}

// NewAlibabaWdkOrderRefundListRequest 初始化AlibabaWdkOrderRefundListAPIRequest对象
func NewAlibabaWdkOrderRefundListRequest() *AlibabaWdkOrderRefundListAPIRequest {
	return &AlibabaWdkOrderRefundListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderRefundListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderRefundListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BatchQueryRefundRequest Setter
// 查询条件
func (r *AlibabaWdkOrderRefundListAPIRequest) SetBatchQueryRefundRequest(_batchQueryRefundRequest *BatchQueryRefundRequest) error {
	r._batchQueryRefundRequest = _batchQueryRefundRequest
	r.Set("batch_query_refund_request", _batchQueryRefundRequest)
	return nil
}

// Get BatchQueryRefundRequest Getter
func (r AlibabaWdkOrderRefundListAPIRequest) GetBatchQueryRefundRequest() *BatchQueryRefundRequest {
	return r._batchQueryRefundRequest
}
