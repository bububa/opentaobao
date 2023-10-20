package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderRefundListAPIRequest) Reset() {
	r._batchQueryRefundRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderRefundListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderRefundListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderRefundListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchQueryRefundRequest is BatchQueryRefundRequest Setter
// 查询条件
func (r *AlibabaWdkOrderRefundListAPIRequest) SetBatchQueryRefundRequest(_batchQueryRefundRequest *BatchQueryRefundRequest) error {
	r._batchQueryRefundRequest = _batchQueryRefundRequest
	r.Set("batch_query_refund_request", _batchQueryRefundRequest)
	return nil
}

// GetBatchQueryRefundRequest BatchQueryRefundRequest Getter
func (r AlibabaWdkOrderRefundListAPIRequest) GetBatchQueryRefundRequest() *BatchQueryRefundRequest {
	return r._batchQueryRefundRequest
}

var poolAlibabaWdkOrderRefundListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderRefundListRequest()
	},
}

// GetAlibabaWdkOrderRefundListRequest 从 sync.Pool 获取 AlibabaWdkOrderRefundListAPIRequest
func GetAlibabaWdkOrderRefundListAPIRequest() *AlibabaWdkOrderRefundListAPIRequest {
	return poolAlibabaWdkOrderRefundListAPIRequest.Get().(*AlibabaWdkOrderRefundListAPIRequest)
}

// ReleaseAlibabaWdkOrderRefundListAPIRequest 将 AlibabaWdkOrderRefundListAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderRefundListAPIRequest(v *AlibabaWdkOrderRefundListAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderRefundListAPIRequest.Put(v)
}
