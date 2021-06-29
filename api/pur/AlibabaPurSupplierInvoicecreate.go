package pur

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/pur"
)

/* 
preInvoice创建 
alibaba.pur.supplier.invoicecreate

preInvoice创建
*/
func AlibabaPurSupplierInvoicecreate(clt *core.SDKClient, req *pur.AlibabaPurSupplierInvoicecreateRequest, session string) (*pur.AlibabaPurSupplierInvoicecreateAPIResponse, error) {
    var resp pur.AlibabaPurSupplierInvoicecreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
