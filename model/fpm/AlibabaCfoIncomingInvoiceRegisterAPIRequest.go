package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfoincominginvoiceregisterAPIRequest 发票登记接口 API请求
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
type AlibabacfoincominginvoiceregisterAPIRequest struct {
	model.Params
	// 发票登记请求体
	_invoiceRegisterRequest *InvoiceRegisterRequest
}

// NewAlibabacfoincominginvoiceregisterRequest 初始化AlibabacfoincominginvoiceregisterAPIRequest对象
func NewAlibabacfoincominginvoiceregisterRequest() *AlibabacfoincominginvoiceregisterAPIRequest {
	return &AlibabacfoincominginvoiceregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacfoincominginvoiceregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.cfo.incoming.invoice.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacfoincominginvoiceregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacfoincominginvoiceregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceRegisterRequest is InvoiceRegisterRequest Setter
// 发票登记请求体
func (r *AlibabacfoincominginvoiceregisterAPIRequest) SetInvoiceRegisterRequest(_invoiceRegisterRequest *InvoiceRegisterRequest) error {
	r._invoiceRegisterRequest = _invoiceRegisterRequest
	r.Set("invoice_register_request", _invoiceRegisterRequest)
	return nil
}

// GetInvoiceRegisterRequest InvoiceRegisterRequest Getter
func (r AlibabacfoincominginvoiceregisterAPIRequest) GetInvoiceRegisterRequest() *InvoiceRegisterRequest {
	return r._invoiceRegisterRequest
}
