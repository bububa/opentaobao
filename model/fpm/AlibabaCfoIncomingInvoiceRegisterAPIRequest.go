package fpm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
