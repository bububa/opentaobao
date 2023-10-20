package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierInvoicecreateAPIRequest preInvoice创建 API请求
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
type AlibabaPurSupplierInvoicecreateAPIRequest struct {
	model.Params
	// 预发票头信息
	_invoice *SupplierPreInvoiceInfoVo
}

// NewAlibabaPurSupplierInvoicecreateRequest 初始化AlibabaPurSupplierInvoicecreateAPIRequest对象
func NewAlibabaPurSupplierInvoicecreateRequest() *AlibabaPurSupplierInvoicecreateAPIRequest {
	return &AlibabaPurSupplierInvoicecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierInvoicecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.invoicecreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierInvoicecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurSupplierInvoicecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoice is Invoice Setter
// 预发票头信息
func (r *AlibabaPurSupplierInvoicecreateAPIRequest) SetInvoice(_invoice *SupplierPreInvoiceInfoVo) error {
	r._invoice = _invoice
	r.Set("invoice", _invoice)
	return nil
}

// GetInvoice Invoice Getter
func (r AlibabaPurSupplierInvoicecreateAPIRequest) GetInvoice() *SupplierPreInvoiceInfoVo {
	return r._invoice
}
