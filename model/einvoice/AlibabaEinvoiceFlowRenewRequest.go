package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单(入驻、加盘、续约)续约 APIRequest
alibaba.einvoice.flow.renew

工单(含入驻、加盘、续约工单)续约能力开放
*/
type AlibabaEinvoiceFlowRenewRequest struct {
    model.Params

    // 续约请求参数
    invoiceFlowRenewDto   *InvoiceFlowRenewDto 

}

func NewAlibabaEinvoiceFlowRenewRequest() *AlibabaEinvoiceFlowRenewRequest{
    return &AlibabaEinvoiceFlowRenewRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceFlowRenewRequest) GetApiMethodName() string {
    return "alibaba.einvoice.flow.renew"
}

func (r AlibabaEinvoiceFlowRenewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceFlowRenewRequest) SetInvoiceFlowRenewDto(invoiceFlowRenewDto *InvoiceFlowRenewDto) error {
    r.invoiceFlowRenewDto = invoiceFlowRenewDto
    r.Set("invoice_flow_renew_dto", invoiceFlowRenewDto)
    return nil
}

func (r AlibabaEinvoiceFlowRenewRequest) GetInvoiceFlowRenewDto() *InvoiceFlowRenewDto {
    return r.invoiceFlowRenewDto
}

