package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票回调接口 API请求
taobao.bus.invoice.return

汽车票发票回调接口
*/
type TaobaoBusInvoiceReturnRequest struct {
    model.Params
    // 入参对象
    _invoiceParam   *ReceiptDo
}

// 初始化TaobaoBusInvoiceReturnRequest对象
func NewTaobaoBusInvoiceReturnRequest() *TaobaoBusInvoiceReturnRequest{
    return &TaobaoBusInvoiceReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusInvoiceReturnRequest) GetApiMethodName() string {
    return "taobao.bus.invoice.return"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusInvoiceReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceParam Setter
// 入参对象
func (r *TaobaoBusInvoiceReturnRequest) SetInvoiceParam(_invoiceParam *ReceiptDo) error {
    r._invoiceParam = _invoiceParam
    r.Set("invoice_param", _invoiceParam)
    return nil
}

// InvoiceParam Getter
func (r TaobaoBusInvoiceReturnRequest) GetInvoiceParam() *ReceiptDo {
    return r._invoiceParam
}
