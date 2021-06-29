package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
preInvoice创建 APIResponse
alibaba.pur.supplier.invoicecreate

preInvoice创建
*/
type AlibabaPurSupplierInvoicecreateAPIResponse struct {
    model.CommonResponse
    AlibabaPurSupplierInvoicecreateResponse
}

type AlibabaPurSupplierInvoicecreateResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_supplier_invoicecreate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取url的出参
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
