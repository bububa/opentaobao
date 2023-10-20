package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyorderreceiptqueryAPIRequest 订单小票查询 API请求
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
type AlibabatclsaelophyorderreceiptqueryAPIRequest struct {
	model.Params
	// 小票查询请求
	_receiptQueryRequest *ReceiptQueryRequest
}

// NewAlibabatclsaelophyorderreceiptqueryRequest 初始化AlibabatclsaelophyorderreceiptqueryAPIRequest对象
func NewAlibabatclsaelophyorderreceiptqueryRequest() *AlibabatclsaelophyorderreceiptqueryAPIRequest {
	return &AlibabatclsaelophyorderreceiptqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyorderreceiptqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.order.receipt.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyorderreceiptqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyorderreceiptqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptQueryRequest is ReceiptQueryRequest Setter
// 小票查询请求
func (r *AlibabatclsaelophyorderreceiptqueryAPIRequest) SetReceiptQueryRequest(_receiptQueryRequest *ReceiptQueryRequest) error {
	r._receiptQueryRequest = _receiptQueryRequest
	r.Set("receipt_query_request", _receiptQueryRequest)
	return nil
}

// GetReceiptQueryRequest ReceiptQueryRequest Getter
func (r AlibabatclsaelophyorderreceiptqueryAPIRequest) GetReceiptQueryRequest() *ReceiptQueryRequest {
	return r._receiptQueryRequest
}
