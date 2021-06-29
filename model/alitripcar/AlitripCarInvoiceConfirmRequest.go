package alitripcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票确认接口 API请求
alitrip.car.invoice.confirm

飞猪发票回调接口
*/
type AlitripCarInvoiceConfirmRequest struct {
    model.Params
    // 入参对象
    receiptDo   *ReceiptDo
}

// 初始化AlitripCarInvoiceConfirmRequest对象
func NewAlitripCarInvoiceConfirmRequest() *AlitripCarInvoiceConfirmRequest{
    return &AlitripCarInvoiceConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripCarInvoiceConfirmRequest) GetApiMethodName() string {
    return "alitrip.car.invoice.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlitripCarInvoiceConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptDo Setter
// 入参对象
func (r *AlitripCarInvoiceConfirmRequest) SetReceiptDo(receiptDo *ReceiptDo) error {
    r.receiptDo = receiptDo
    r.Set("receipt_do", receiptDo)
    return nil
}

// ReceiptDo Getter
func (r AlitripCarInvoiceConfirmRequest) GetReceiptDo() *ReceiptDo {
    return r.receiptDo
}
