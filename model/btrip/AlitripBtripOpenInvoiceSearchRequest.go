package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
差旅申请用户搜索可用发票列表 API请求
alitrip.btrip.open.invoice.search

差旅申请用户搜索可用发票列表
*/
type AlitripBtripOpenInvoiceSearchRequest struct {
    model.Params
    // 入参对象
    rq   *OpenInvoiceRq
}

// 初始化AlitripBtripOpenInvoiceSearchRequest对象
func NewAlitripBtripOpenInvoiceSearchRequest() *AlitripBtripOpenInvoiceSearchRequest{
    return &AlitripBtripOpenInvoiceSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenInvoiceSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.invoice.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenInvoiceSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenInvoiceSearchRequest) SetRq(rq *OpenInvoiceRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenInvoiceSearchRequest) GetRq() *OpenInvoiceRq {
    return r.rq
}
