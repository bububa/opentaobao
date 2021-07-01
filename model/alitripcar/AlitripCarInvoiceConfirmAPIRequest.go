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
type AlitripCarInvoiceConfirmAPIRequest struct {
    model.Params
    // 入参对象
    _receiptDo   *ReceiptDo
}

// 初始化AlitripCarInvoiceConfirmAPIRequest对象
func NewAlitripCarInvoiceConfirmRequest() *AlitripCarInvoiceConfirmAPIRequest{
    return &AlitripCarInvoiceConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripCarInvoiceConfirmAPIRequest) GetApiMethodName() string {
    return "alitrip.car.invoice.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlitripCarInvoiceConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptDo Setter
// 入参对象
func (r *AlitripCarInvoiceConfirmAPIRequest) SetReceiptDo(_receiptDo *ReceiptDo) error {
    r._receiptDo = _receiptDo
    r.Set("receipt_do", _receiptDo)
    return nil
}

// ReceiptDo Getter
func (r AlitripCarInvoiceConfirmAPIRequest) GetReceiptDo() *ReceiptDo {
    return r._receiptDo
}
