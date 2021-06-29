package alitripreceipt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票查询 APIRequest
alitrip.receipt.seller.invoice.query

飞猪发票查询
*/
type AlitripReceiptSellerInvoiceQueryRequest struct {
    model.Params

    // 入参对象
    queryReceiptParam   *QueryReceiptParam 

}

func NewAlitripReceiptSellerInvoiceQueryRequest() *AlitripReceiptSellerInvoiceQueryRequest{
    return &AlitripReceiptSellerInvoiceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripReceiptSellerInvoiceQueryRequest) GetApiMethodName() string {
    return "alitrip.receipt.seller.invoice.query"
}

func (r AlitripReceiptSellerInvoiceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripReceiptSellerInvoiceQueryRequest) SetQueryReceiptParam(queryReceiptParam *QueryReceiptParam) error {
    r.queryReceiptParam = queryReceiptParam
    r.Set("query_receipt_param", queryReceiptParam)
    return nil
}

func (r AlitripReceiptSellerInvoiceQueryRequest) GetQueryReceiptParam() *QueryReceiptParam {
    return r.queryReceiptParam
}

