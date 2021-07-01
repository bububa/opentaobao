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
type TaobaoBusInvoiceReturnAPIRequest struct {
    model.Params
    // 入参对象
    _invoiceParam   *ReceiptDO
}

// 初始化TaobaoBusInvoiceReturnAPIRequest对象
func NewTaobaoBusInvoiceReturnRequest() *TaobaoBusInvoiceReturnAPIRequest{
    return &TaobaoBusInvoiceReturnAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusInvoiceReturnAPIRequest) GetApiMethodName() string {
    return "taobao.bus.invoice.return"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusInvoiceReturnAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceParam Setter
// 入参对象
func (r *TaobaoBusInvoiceReturnAPIRequest) SetInvoiceParam(_invoiceParam *ReceiptDO) error {
    r._invoiceParam = _invoiceParam
    r.Set("invoice_param", _invoiceParam)
    return nil
}

// InvoiceParam Getter
func (r TaobaoBusInvoiceReturnAPIRequest) GetInvoiceParam() *ReceiptDO {
    return r._invoiceParam
}
