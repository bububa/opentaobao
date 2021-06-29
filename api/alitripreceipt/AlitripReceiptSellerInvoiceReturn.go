package alitripreceipt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripreceipt"
)

/* 
飞猪发票商家回调接口 
alitrip.receipt.seller.invoice.return

飞猪发票回调接口
*/
func AlitripReceiptSellerInvoiceReturn(clt *core.SDKClient, req *alitripreceipt.AlitripReceiptSellerInvoiceReturnRequest, session string) (*alitripreceipt.AlitripReceiptSellerInvoiceReturnAPIResponse, error) {
    var resp alitripreceipt.AlitripReceiptSellerInvoiceReturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
