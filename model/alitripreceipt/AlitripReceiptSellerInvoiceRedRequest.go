package alitripreceipt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票冲红 APIRequest
alitrip.receipt.seller.invoice.red

飞猪发票创建
*/
type AlitripReceiptSellerInvoiceRedRequest struct {
    model.Params

    // 入参对象
    redReceiptParam   *RedReceiptParam 

}

func NewAlitripReceiptSellerInvoiceRedRequest() *AlitripReceiptSellerInvoiceRedRequest{
    return &AlitripReceiptSellerInvoiceRedRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripReceiptSellerInvoiceRedRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.red"
}

func (r AlitripReceiptSellerInvoiceRedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripReceiptSellerInvoiceRedRequest) SetRedReceiptParam(redReceiptParam *RedReceiptParam) error {
    r.redReceiptParam = redReceiptParam
    r.Set("red_receipt_param", redReceiptParam)
    return nil
}

func (r AlitripReceiptSellerInvoiceRedRequest) GetRedReceiptParam() *RedReceiptParam {
    return r.redReceiptParam
}

