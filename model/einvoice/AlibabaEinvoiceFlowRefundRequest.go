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
type AlibabaEinvoiceFlowRefundRequest struct {
    model.Params
    // 退订请求参数
    _invoiceFlowRefund   *InvoiceFlowRefundDTO
}

// 初始化AlibabaEinvoiceFlowRefundRequest对象
func NewAlibabaEinvoiceFlowRefundRequest() *AlibabaEinvoiceFlowRefundRequest{
    return &AlibabaEinvoiceFlowRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowRefundRequest) GetApiMethodName() string {
    return "alibaba.einvoice.flow.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceFlowRefund Setter
// 退订请求参数
func (r *AlibabaEinvoiceFlowRefundRequest) SetInvoiceFlowRefund(_invoiceFlowRefund *InvoiceFlowRefundDTO) error {
    r._invoiceFlowRefund = _invoiceFlowRefund
    r.Set("invoice_flow_refund", _invoiceFlowRefund)
    return nil
}

// InvoiceFlowRefund Getter
func (r AlibabaEinvoiceFlowRefundRequest) GetInvoiceFlowRefund() *InvoiceFlowRefundDTO {
    return r._invoiceFlowRefund
}
