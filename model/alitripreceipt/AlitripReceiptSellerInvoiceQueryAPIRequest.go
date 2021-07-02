package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptSellerInvoiceQueryAPIRequest 飞猪发票查询 API请求
// alitrip.receipt.seller.invoice.query
//
// 飞猪发票查询
type AlitripReceiptSellerInvoiceQueryAPIRequest struct {
	model.Params
	// 入参对象
	_queryReceiptParam *QueryReceiptParam
}

// NewAlitripReceiptSellerInvoiceQueryRequest 初始化AlitripReceiptSellerInvoiceQueryAPIRequest对象
func NewAlitripReceiptSellerInvoiceQueryRequest() *AlitripReceiptSellerInvoiceQueryAPIRequest {
	return &AlitripReceiptSellerInvoiceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.seller.invoice.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryReceiptParam is QueryReceiptParam Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceQueryAPIRequest) SetQueryReceiptParam(_queryReceiptParam *QueryReceiptParam) error {
	r._queryReceiptParam = _queryReceiptParam
	r.Set("query_receipt_param", _queryReceiptParam)
	return nil
}

// GetQueryReceiptParam QueryReceiptParam Getter
func (r AlitripReceiptSellerInvoiceQueryAPIRequest) GetQueryReceiptParam() *QueryReceiptParam {
	return r._queryReceiptParam
}
