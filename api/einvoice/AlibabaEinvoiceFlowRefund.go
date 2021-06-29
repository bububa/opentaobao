package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
退订工单(入驻、加盘、续约) 
alibaba.einvoice.flow.refund

电子发票工单系统，工单退订能力开放
*/
func AlibabaEinvoiceFlowRefund(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceFlowRefundRequest, session string) (*einvoice.AlibabaEinvoiceFlowRefundAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceFlowRefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
