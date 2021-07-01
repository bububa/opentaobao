package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
开票量核对接口 
alibaba.einvoice.amount.check

跟开票服务商核对历史开票量，用来对账
*/
func AlibabaEinvoiceAmountCheck(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceAmountCheckAPIRequest, session string) (*einvoice.AlibabaEinvoiceAmountCheckAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceAmountCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
