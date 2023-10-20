package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyOrderReceiptQueryAPIRequest 订单小票查询 API请求
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
type AlibabaTclsAelophyOrderReceiptQueryAPIRequest struct {
	model.Params
	// 小票查询请求
	_receiptQueryRequest *ReceiptQueryRequest
}

// NewAlibabaTclsAelophyOrderReceiptQueryRequest 初始化AlibabaTclsAelophyOrderReceiptQueryAPIRequest对象
func NewAlibabaTclsAelophyOrderReceiptQueryRequest() *AlibabaTclsAelophyOrderReceiptQueryAPIRequest {
	return &AlibabaTclsAelophyOrderReceiptQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyOrderReceiptQueryAPIRequest) Reset() {
	r._receiptQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.order.receipt.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptQueryRequest is ReceiptQueryRequest Setter
// 小票查询请求
func (r *AlibabaTclsAelophyOrderReceiptQueryAPIRequest) SetReceiptQueryRequest(_receiptQueryRequest *ReceiptQueryRequest) error {
	r._receiptQueryRequest = _receiptQueryRequest
	r.Set("receipt_query_request", _receiptQueryRequest)
	return nil
}

// GetReceiptQueryRequest ReceiptQueryRequest Getter
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetReceiptQueryRequest() *ReceiptQueryRequest {
	return r._receiptQueryRequest
}

var poolAlibabaTclsAelophyOrderReceiptQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyOrderReceiptQueryRequest()
	},
}

// GetAlibabaTclsAelophyOrderReceiptQueryRequest 从 sync.Pool 获取 AlibabaTclsAelophyOrderReceiptQueryAPIRequest
func GetAlibabaTclsAelophyOrderReceiptQueryAPIRequest() *AlibabaTclsAelophyOrderReceiptQueryAPIRequest {
	return poolAlibabaTclsAelophyOrderReceiptQueryAPIRequest.Get().(*AlibabaTclsAelophyOrderReceiptQueryAPIRequest)
}

// ReleaseAlibabaTclsAelophyOrderReceiptQueryAPIRequest 将 AlibabaTclsAelophyOrderReceiptQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyOrderReceiptQueryAPIRequest(v *AlibabaTclsAelophyOrderReceiptQueryAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyOrderReceiptQueryAPIRequest.Put(v)
}
