package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurSupplierInvoicecreate preInvoice创建
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
func AlibabaPurSupplierInvoicecreate(clt *core.SDKClient, req *pur.AlibabaPurSupplierInvoicecreateAPIRequest, resp *pur.AlibabaPurSupplierInvoicecreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
