package alitripcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票确认接口 APIRequest
alitrip.car.invoice.confirm

飞猪发票回调接口
*/
type AlitripCarInvoiceConfirmRequest struct {
    model.Params

    // 入参对象
    receiptDo   *ReceiptDo 

}

func NewAlitripCarInvoiceConfirmRequest() *AlitripCarInvoiceConfirmRequest{
    return &AlitripCarInvoiceConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripCarInvoiceConfirmRequest) GetApiMethodName() string {
    return "alitrip.car.invoice.confirm"
}

func (r AlitripCarInvoiceConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripCarInvoiceConfirmRequest) SetReceiptDo(receiptDo *ReceiptDo) error {
    r.receiptDo = receiptDo
    r.Set("receipt_do", receiptDo)
    return nil
}

func (r AlitripCarInvoiceConfirmRequest) GetReceiptDo() *ReceiptDo {
    return r.receiptDo
}

