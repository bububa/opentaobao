package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
回传订单退款审核结果 
alibaba.einvoice.order.refund.update

ISV回传订单退款审核结果
*/
func AlibabaEinvoiceOrderRefundUpdate(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceOrderRefundUpdateRequest, session string) (*einvoice.AlibabaEinvoiceOrderRefundUpdateAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceOrderRefundUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
