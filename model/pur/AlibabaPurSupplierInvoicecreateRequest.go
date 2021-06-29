package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
preInvoice创建 APIRequest
alibaba.pur.supplier.invoicecreate

preInvoice创建
*/
type AlibabaPurSupplierInvoicecreateRequest struct {
    model.Params

    // 预发票头信息
    invoice   *SupplierPreInvoiceInfoVO 

}

func NewAlibabaPurSupplierInvoicecreateRequest() *AlibabaPurSupplierInvoicecreateRequest{
    return &AlibabaPurSupplierInvoicecreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurSupplierInvoicecreateRequest) GetApiMethodName() string {
    return "alibaba.pur.supplier.invoicecreate"
}

func (r AlibabaPurSupplierInvoicecreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurSupplierInvoicecreateRequest) SetInvoice(invoice *SupplierPreInvoiceInfoVO) error {
    r.invoice = invoice
    r.Set("invoice", invoice)
    return nil
}

func (r AlibabaPurSupplierInvoicecreateRequest) GetInvoice() *SupplierPreInvoiceInfoVO {
    return r.invoice
}

