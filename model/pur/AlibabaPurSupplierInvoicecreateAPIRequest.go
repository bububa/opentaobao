package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapursupplierinvoicecreateAPIRequest preInvoice创建 API请求
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
type AlibabapursupplierinvoicecreateAPIRequest struct {
	model.Params
	// 预发票头信息
	_invoice *SupplierPreInvoiceInfoVo
}

// NewAlibabapursupplierinvoicecreateRequest 初始化AlibabapursupplierinvoicecreateAPIRequest对象
func NewAlibabapursupplierinvoicecreateRequest() *AlibabapursupplierinvoicecreateAPIRequest {
	return &AlibabapursupplierinvoicecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapursupplierinvoicecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.invoicecreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapursupplierinvoicecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapursupplierinvoicecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoice is Invoice Setter
// 预发票头信息
func (r *AlibabapursupplierinvoicecreateAPIRequest) SetInvoice(_invoice *SupplierPreInvoiceInfoVo) error {
	r._invoice = _invoice
	r.Set("invoice", _invoice)
	return nil
}

// GetInvoice Invoice Getter
func (r AlibabapursupplierinvoicecreateAPIRequest) GetInvoice() *SupplierPreInvoiceInfoVo {
	return r._invoice
}
