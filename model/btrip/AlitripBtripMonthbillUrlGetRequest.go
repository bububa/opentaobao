package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
月账单数据查询 APIRequest
alitrip.btrip.monthbill.url.get

月账单数据查询
*/
type AlitripBtripMonthbillUrlGetRequest struct {
    model.Params

    // 请求对象
    rq   *OpenAccountRq 

}

func NewAlitripBtripMonthbillUrlGetRequest() *AlitripBtripMonthbillUrlGetRequest{
    return &AlitripBtripMonthbillUrlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripMonthbillUrlGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.monthbill.url.get"
}

func (r AlitripBtripMonthbillUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripMonthbillUrlGetRequest) SetRq(rq *OpenAccountRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripMonthbillUrlGetRequest) GetRq() *OpenAccountRq {
    return r.rq
}

