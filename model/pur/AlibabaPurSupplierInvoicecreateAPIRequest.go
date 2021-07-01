package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurSupplierInvoicecreateAPIRequest
preInvoice创建 API请求
alibaba.pur.supplier.invoicecreate

preInvoice创建 */
type AlibabaPurSupplierInvoicecreateAPIRequest struct {
	model.Params
	// 预发票头信息
	_invoice *SupplierPreInvoiceInfoVO
}

// New
