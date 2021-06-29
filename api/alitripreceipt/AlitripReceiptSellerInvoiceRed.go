package alitripreceipt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripreceipt"
)

/* 
飞猪发票冲红 
alitrip.receipt.seller.invoice.red

飞猪发票创建
*/
func AlitripReceiptSellerInvoiceRed(clt *core.SDKClient, req *alitripreceipt.AlitripReceiptSellerInvoiceRedRequest, session string) (*alitripreceipt.AlitripReceiptSellerInvoiceRedAPIResponse, error) {
    var resp alitripreceipt.AlitripReceiptSellerInvoiceRedAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}