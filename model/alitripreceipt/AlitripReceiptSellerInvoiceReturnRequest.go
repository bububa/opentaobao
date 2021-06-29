package alitripreceipt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票商家回调接口 APIRequest
alitrip.receipt.seller.invoice.return

飞猪发票回调接口
*/
type AlitripReceiptSellerInvoiceReturnRequest struct {
    model.Params

    // 入参对象
    receiptDo   *ReceiptDo 

}

func NewAlitripReceiptSellerInvoiceReturnRequest() *AlitripReceiptSellerInvoiceReturnRequest{
    return &AlitripReceiptSellerInvoiceReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripReceiptSellerInvoiceReturnRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.return"
}

func (r AlitripReceiptSellerInvoiceReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripReceiptSellerInvoiceReturnRequest) SetReceiptDo(receiptDo *ReceiptDo) error {
    r.receiptDo = receiptDo
    r.Set("receipt_do", receiptDo)
    return nil
}

func (r AlitripReceiptSellerInvoiceReturnRequest) GetReceiptDo() *ReceiptDo {
    return r.receiptDo
}

