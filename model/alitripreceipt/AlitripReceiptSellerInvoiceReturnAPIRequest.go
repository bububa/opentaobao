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
type AlitripReceiptSellerInvoiceReturnAPIRequest struct {
    model.Params
    // 入参对象
    _receiptDo   *ReceiptDo
}

// 初始化AlitripReceiptSellerInvoiceReturnAPIRequest对象
func NewAlitripReceiptSellerInvoiceReturnRequest() *AlitripReceiptSellerInvoiceReturnAPIRequest{
    return &AlitripReceiptSellerInvoiceReturnAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.return"
}

// IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptDo Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceReturnAPIRequest) SetReceiptDo(_receiptDo *ReceiptDo) error {
    r._receiptDo = _receiptDo
    r.Set("receipt_do", _receiptDo)
    return nil
}

// ReceiptDo Getter
func (r AlitripReceiptSellerInvoiceReturnAPIRequest) GetReceiptDo() *ReceiptDo {
    return r._receiptDo
}
