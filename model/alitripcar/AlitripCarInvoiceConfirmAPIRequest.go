package alitripcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripCarInvoiceConfirmAPIRequest 发票确认接口 API请求
// alitrip.car.invoice.confirm
//
// 飞猪发票回调接口
type AlitripCarInvoiceConfirmAPIRequest struct {
	model.Params
	// 入参对象
	_receiptDo *ReceiptDo
}

// NewAlitripCarInvoiceConfirmRequest 初始化AlitripCarInvoiceConfirmAPIRequest对象
func NewAlitripCarInvoiceConfirmRequest() *AlitripCarInvoiceConfirmAPIRequest {
	return &AlitripCarInvoiceConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripCarInvoiceConfirmAPIRequest) Reset() {
	r._receiptDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripCarInvoiceConfirmAPIRequest) GetApiMethodName() string {
	return "alitrip.car.invoice.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripCarInvoiceConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripCarInvoiceConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptDo is ReceiptDo Setter
// 入参对象
func (r *AlitripCarInvoiceConfirmAPIRequest) SetReceiptDo(_receiptDo *ReceiptDo) error {
	r._receiptDo = _receiptDo
	r.Set("receipt_do", _receiptDo)
	return nil
}

// GetReceiptDo ReceiptDo Getter
func (r AlitripCarInvoiceConfirmAPIRequest) GetReceiptDo() *ReceiptDo {
	return r._receiptDo
}

var poolAlitripCarInvoiceConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripCarInvoiceConfirmRequest()
	},
}

// GetAlitripCarInvoiceConfirmRequest 从 sync.Pool 获取 AlitripCarInvoiceConfirmAPIRequest
func GetAlitripCarInvoiceConfirmAPIRequest() *AlitripCarInvoiceConfirmAPIRequest {
	return poolAlitripCarInvoiceConfirmAPIRequest.Get().(*AlitripCarInvoiceConfirmAPIRequest)
}

// ReleaseAlitripCarInvoiceConfirmAPIRequest 将 AlitripCarInvoiceConfirmAPIRequest 放入 sync.Pool
func ReleaseAlitripCarInvoiceConfirmAPIRequest(v *AlitripCarInvoiceConfirmAPIRequest) {
	v.Reset()
	poolAlitripCarInvoiceConfirmAPIRequest.Put(v)
}
