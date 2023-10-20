package fpm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoiceRegisterAPIRequest 发票登记接口 API请求
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
type AlibabaCfoIncomingInvoiceRegisterAPIRequest struct {
	model.Params
	// 发票登记请求体
	_invoiceRegisterRequest *InvoiceRegisterRequest
}

// NewAlibabaCfoIncomingInvoiceRegisterRequest 初始化AlibabaCfoIncomingInvoiceRegisterAPIRequest对象
func NewAlibabaCfoIncomingInvoiceRegisterRequest() *AlibabaCfoIncomingInvoiceRegisterAPIRequest {
	return &AlibabaCfoIncomingInvoiceRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCfoIncomingInvoiceRegisterAPIRequest) Reset() {
	r._invoiceRegisterRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCfoIncomingInvoiceRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.cfo.incoming.invoice.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCfoIncomingInvoiceRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCfoIncomingInvoiceRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceRegisterRequest is InvoiceRegisterRequest Setter
// 发票登记请求体
func (r *AlibabaCfoIncomingInvoiceRegisterAPIRequest) SetInvoiceRegisterRequest(_invoiceRegisterRequest *InvoiceRegisterRequest) error {
	r._invoiceRegisterRequest = _invoiceRegisterRequest
	r.Set("invoice_register_request", _invoiceRegisterRequest)
	return nil
}

// GetInvoiceRegisterRequest InvoiceRegisterRequest Getter
func (r AlibabaCfoIncomingInvoiceRegisterAPIRequest) GetInvoiceRegisterRequest() *InvoiceRegisterRequest {
	return r._invoiceRegisterRequest
}

var poolAlibabaCfoIncomingInvoiceRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCfoIncomingInvoiceRegisterRequest()
	},
}

// GetAlibabaCfoIncomingInvoiceRegisterRequest 从 sync.Pool 获取 AlibabaCfoIncomingInvoiceRegisterAPIRequest
func GetAlibabaCfoIncomingInvoiceRegisterAPIRequest() *AlibabaCfoIncomingInvoiceRegisterAPIRequest {
	return poolAlibabaCfoIncomingInvoiceRegisterAPIRequest.Get().(*AlibabaCfoIncomingInvoiceRegisterAPIRequest)
}

// ReleaseAlibabaCfoIncomingInvoiceRegisterAPIRequest 将 AlibabaCfoIncomingInvoiceRegisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaCfoIncomingInvoiceRegisterAPIRequest(v *AlibabaCfoIncomingInvoiceRegisterAPIRequest) {
	v.Reset()
	poolAlibabaCfoIncomingInvoiceRegisterAPIRequest.Put(v)
}
