package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
preInvoice创建 API请求
alibaba.pur.supplier.invoicecreate

preInvoice创建
*/
type AlibabaPurSupplierInvoicecreateRequest struct {
    model.Params
    // 预发票头信息
    invoice   *SupplierPreInvoiceInfoVO
}

// 初始化AlibabaPurSupplierInvoicecreateRequest对象
func NewAlibabaPurSupplierInvoicecreateRequest() *AlibabaPurSupplierInvoicecreateRequest{
    return &AlibabaPurSupplierInvoicecreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierInvoicecreateRequest) GetApiMethodName() string {
    return "alibaba.pur.supplier.invoicecreate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierInvoicecreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Invoice Setter
// 预发票头信息
func (r *AlibabaPurSupplierInvoicecreateRequest) SetInvoice(invoice *SupplierPreInvoiceInfoVO) error {
    r.invoice = invoice
    r.Set("invoice", invoice)
    return nil
}

// Invoice Getter
func (r AlibabaPurSupplierInvoicecreateRequest) GetInvoice() *SupplierPreInvoiceInfoVO {
    return r.invoice
}
