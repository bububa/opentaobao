package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票回调接口 APIRequest
taobao.bus.invoice.return

汽车票发票回调接口
*/
type TaobaoBusInvoiceReturnRequest struct {
    model.Params

    // 入参对象
    invoiceParam   *ReceiptDo 

}

func NewTaobaoBusInvoiceReturnRequest() *TaobaoBusInvoiceReturnRequest{
    return &TaobaoBusInvoiceReturnRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusInvoiceReturnRequest) GetApiMethodName() string {
    return "taobao.bus.invoice.return"
}

func (r TaobaoBusInvoiceReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusInvoiceReturnRequest) SetInvoiceParam(invoiceParam *ReceiptDo) error {
    r.invoiceParam = invoiceParam
    r.Set("invoice_param", invoiceParam)
    return nil
}

func (r TaobaoBusInvoiceReturnRequest) GetInvoiceParam() *ReceiptDo {
    return r.invoiceParam
}

