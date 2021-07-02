package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptSellerInvoiceRedAPIRequest 飞猪发票冲红 API请求
// alitrip.receipt.seller.invoice.red
//
// 飞猪发票创建
type AlitripReceiptSellerInvoiceRedAPIRequest struct {
	model.Params
	// 入参对象
	_redReceiptParam *RedReceiptParam
}

// NewAlitripReceiptSellerInvoiceRedRequest 初始化AlitripReceiptSellerInvoiceRedAPIRequest对象
func NewAlitripReceiptSellerInvoiceRedRequest() *AlitripReceiptSellerInvoiceRedAPIRequest {
	return &AlitripReceiptSellerInvoiceRedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceRedAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.seller.invoice.red"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceRedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RedReceiptParam Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceRedAPIRequest) SetRedReceiptParam(_redReceiptParam *RedReceiptParam) error {
	r._redReceiptParam = _redReceiptParam
	r.Set("red_receipt_param", _redReceiptParam)
	return nil
}

// Get RedReceiptParam Getter
func (r AlitripReceiptSellerInvoiceRedAPIRequest) GetRedReceiptParam() *RedReceiptParam {
	return r._redReceiptParam
}
