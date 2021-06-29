package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发票申请 APIRequest
alibaba.einvoice.prod.apply

提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
*/
type AlibabaEinvoiceProdApplyRequest struct {
    model.Params

    // 申请开票请求
    paramInvoiceApplyDto   *InvoiceApplyDto 

}

func NewAlibabaEinvoiceProdApplyRequest() *AlibabaEinvoiceProdApplyRequest{
    return &AlibabaEinvoiceProdApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceProdApplyRequest) GetApiMethodName() string {
    return "alibaba.einvoice.prod.apply"
}

func (r AlibabaEinvoiceProdApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceProdApplyRequest) SetParamInvoiceApplyDto(paramInvoiceApplyDto *InvoiceApplyDto) error {
    r.paramInvoiceApplyDto = paramInvoiceApplyDto
    r.Set("param_invoice_apply_dto", paramInvoiceApplyDto)
    return nil
}

func (r AlibabaEinvoiceProdApplyRequest) GetParamInvoiceApplyDto() *InvoiceApplyDto {
    return r.paramInvoiceApplyDto
}

