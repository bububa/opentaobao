package alitripreceipt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票商家回调接口 API请求
alitrip.receipt.seller.invoice.return

飞猪发票回调接口
*/
type AlitripReceiptSellerInvoiceReturnRequest struct {
    model.Params
    // 入参对象
    _receiptDo   *ReceiptDo
}

// 初始化AlitripReceiptSellerInvoiceReturnRequest对象
func NewAlitripReceiptSellerInvoiceReturnRequest() *AlitripReceiptSellerInvoiceReturnRequest{
    return &AlitripReceiptSellerInvoiceReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceReturnRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.return"
}

// IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptDo Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceReturnRequest) SetReceiptDo(_receiptDo *ReceiptDo) error {
    r._receiptDo = _receiptDo
    r.Set("receipt_do", _receiptDo)
    return nil
}

// ReceiptDo Getter
func (r AlitripReceiptSellerInvoiceReturnRequest) GetReceiptDo() *ReceiptDo {
    return r._receiptDo
}
