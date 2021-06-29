package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
月账单数据查询 API请求
alitrip.btrip.monthbill.url.get

月账单数据查询
*/
type AlitripBtripMonthbillUrlGetRequest struct {
    model.Params
    // 请求对象
    rq   *OpenAccountRq
}

// 初始化AlitripBtripMonthbillUrlGetRequest对象
func NewAlitripBtripMonthbillUrlGetRequest() *AlitripBtripMonthbillUrlGetRequest{
    return &AlitripBtripMonthbillUrlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripMonthbillUrlGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.monthbill.url.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripMonthbillUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripMonthbillUrlGetRequest) SetRq(rq *OpenAccountRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripMonthbillUrlGetRequest) GetRq() *OpenAccountRq {
    return r.rq
}
