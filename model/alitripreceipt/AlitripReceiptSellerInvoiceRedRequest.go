package alitripreceipt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票冲红 API请求
alitrip.receipt.seller.invoice.red

飞猪发票创建
*/
type AlitripReceiptSellerInvoiceRedRequest struct {
    model.Params
    // 入参对象
    redReceiptParam   *RedReceiptParam
}

// 初始化AlitripReceiptSellerInvoiceRedRequest对象
func NewAlitripReceiptSellerInvoiceRedRequest() *AlitripReceiptSellerInvoiceRedRequest{
    return &AlitripReceiptSellerInvoiceRedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceRedRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.red"
}

// IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceRedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RedReceiptParam Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceRedRequest) SetRedReceiptParam(redReceiptParam *RedReceiptParam) error {
    r.redReceiptParam = redReceiptParam
    r.Set("red_receipt_param", redReceiptParam)
    return nil
}

// RedReceiptParam Getter
func (r AlitripReceiptSellerInvoiceRedRequest) GetRedReceiptParam() *RedReceiptParam {
    return r.redReceiptParam
}
