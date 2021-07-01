package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退订工单(入驻、加盘、续约) API请求
alibaba.einvoice.flow.refund

电子发票工单系统，工单退订能力开放
*/
type AlibabaEinvoiceFlowRefundAPIRequest struct {
    model.Params
    // 退订请求参数
    _invoiceFlowRefund   *InvoiceFlowRefundDto
}

// 初始化AlibabaEinvoiceFlowRefundAPIRequest对象
func NewAlibabaEinvoiceFlowRefundRequest() *AlibabaEinvoiceFlowRefundAPIRequest{
    return &AlibabaEinvoiceFlowRefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.flow.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceFlowRefund Setter
// 退订请求参数
func (r *AlibabaEinvoiceFlowRefundAPIRequest) SetInvoiceFlowRefund(_invoiceFlowRefund *InvoiceFlowRefundDto) error {
    r._invoiceFlowRefund = _invoiceFlowRefund
    r.Set("invoice_flow_refund", _invoiceFlowRefund)
    return nil
}

// InvoiceFlowRefund Getter
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetInvoiceFlowRefund() *InvoiceFlowRefundDto {
    return r._invoiceFlowRefund
}
