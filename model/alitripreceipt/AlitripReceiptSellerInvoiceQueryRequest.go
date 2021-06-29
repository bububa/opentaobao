package alitripreceipt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票查询 API请求
alitrip.receipt.seller.invoice.query

飞猪发票查询
*/
type AlitripReceiptSellerInvoiceQueryRequest struct {
    model.Params
    // 入参对象
    _queryReceiptParam   *QueryReceiptParam
}

// 初始化AlitripReceiptSellerInvoiceQueryRequest对象
func NewAlitripReceiptSellerInvoiceQueryRequest() *AlitripReceiptSellerInvoiceQueryRequest{
    return &AlitripReceiptSellerInvoiceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripReceiptSellerInvoiceQueryRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripReceiptSellerInvoiceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryReceiptParam Setter
// 入参对象
func (r *AlitripReceiptSellerInvoiceQueryRequest) SetQueryReceiptParam(_queryReceiptParam *QueryReceiptParam) error {
    r._queryReceiptParam = _queryReceiptParam
    r.Set("query_receipt_param", _queryReceiptParam)
    return nil
}

// QueryReceiptParam Getter
func (r AlitripReceiptSellerInvoiceQueryRequest) GetQueryReceiptParam() *QueryReceiptParam {
    return r._queryReceiptParam
}
