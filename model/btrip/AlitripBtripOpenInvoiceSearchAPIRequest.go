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
type AlitripBtripOpenInvoiceSearchAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenInvoiceRq
}

// 初始化AlitripBtripOpenInvoiceSearchAPIRequest对象
func NewAlitripBtripOpenInvoiceSearchRequest() *AlitripBtripOpenInvoiceSearchAPIRequest{
    return &AlitripBtripOpenInvoiceSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.invoice.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenInvoiceSearchAPIRequest) SetRq(_rq *OpenInvoiceRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetRq() *OpenInvoiceRq {
    return r._rq
}
