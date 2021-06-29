package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
差旅申请用户搜索可用发票列表 APIRequest
alitrip.btrip.open.invoice.search

差旅申请用户搜索可用发票列表
*/
type AlitripBtripOpenInvoiceSearchRequest struct {
    model.Params

    // 入参对象
    rq   *OpenInvoiceRq 

}

func NewAlitripBtripOpenInvoiceSearchRequest() *AlitripBtripOpenInvoiceSearchRequest{
    return &AlitripBtripOpenInvoiceSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenInvoiceSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.invoice.search"
}

func (r AlitripBtripOpenInvoiceSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenInvoiceSearchRequest) SetRq(rq *OpenInvoiceRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenInvoiceSearchRequest) GetRq() *OpenInvoiceRq {
    return r.rq
}

