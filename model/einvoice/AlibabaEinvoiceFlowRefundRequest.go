package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退订工单(入驻、加盘、续约) APIRequest
alibaba.einvoice.flow.refund

电子发票工单系统，工单退订能力开放
*/
type AlibabaEinvoiceFlowRefundRequest struct {
    model.Params

    // 退订请求参数
    invoiceFlowRefund   *InvoiceFlowRefundDto 

}

func NewAlibabaEinvoiceFlowRefundRequest() *AlibabaEinvoiceFlowRefundRequest{
    return &AlibabaEinvoiceFlowRefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceFlowRefundRequest) GetApiMethodName() string {
    return "alibaba.einvoice.flow.refund"
}

func (r AlibabaEinvoiceFlowRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceFlowRefundRequest) SetInvoiceFlowRefund(invoiceFlowRefund *InvoiceFlowRefundDto) error {
    r.invoiceFlowRefund = invoiceFlowRefund
    r.Set("invoice_flow_refund", invoiceFlowRefund)
    return nil
}

func (r AlibabaEinvoiceFlowRefundRequest) GetInvoiceFlowRefund() *InvoiceFlowRefundDto {
    return r.invoiceFlowRefund
}

