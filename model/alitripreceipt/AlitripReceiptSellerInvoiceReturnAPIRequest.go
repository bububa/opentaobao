package alitripreceipt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptSellerInvoiceReturnAPIRequest 飞猪发票商家回调接口 API请求
// alitrip.receipt.seller.invoice.return
//
// 飞猪发票回调接口
type AlitripReceiptSellerInvoiceReturnAPIRequest struct {
	model.Params
	// 入参对象
	_receiptDo *ReceiptDo
}

// NewAlitripReceiptSellerInvoiceReturnRequest 初始化AlitripReceiptSellerInvoiceReturnAPIRequest对象
func NewAlitripReceiptSellerInvoiceReturnRequest() *AlitripReceiptSellerInvoiceReturnAPIRequest {
	return &AlitripReceiptSellerInvoiceReturnAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripReceiptSellerInvoiceReturnAPIRequest) Reset() {
	r._receiptDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.seller.invoice.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptDo is ReceiptDo Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceReturnAPIRequest) SetReceiptDo(_receiptDo *ReceiptDo) error {
	r._receiptDo = _receiptDo
	r.Set("receipt_do", _receiptDo)
	return nil
}

// GetReceiptDo ReceiptDo Getter
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetReceiptDo() *ReceiptDo {
	return r._receiptDo
}

var poolAlitripReceiptSellerInvoiceReturnAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripReceiptSellerInvoiceReturnRequest()
	},
}

// GetAlitripReceiptSellerInvoiceReturnRequest 从 sync.Pool 获取 AlitripReceiptSellerInvoiceReturnAPIRequest
func GetAlitripReceiptSellerInvoiceReturnAPIRequest() *AlitripReceiptSellerInvoiceReturnAPIRequest {
	return poolAlitripReceiptSellerInvoiceReturnAPIRequest.Get().(*AlitripReceiptSellerInvoiceReturnAPIRequest)
}

// ReleaseAlitripReceiptSellerInvoiceReturnAPIRequest 将 AlitripReceiptSellerInvoiceReturnAPIRequest 放入 sync.Pool
func ReleaseAlitripReceiptSellerInvoiceReturnAPIRequest(v *AlitripReceiptSellerInvoiceReturnAPIRequest) {
	v.Reset()
	poolAlitripReceiptSellerInvoiceReturnAPIRequest.Put(v)
}
