package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取机票订单列表 APIRequest
alitrip.btrip.flight.order.search

第三方OA厂商获取机票订单列表
*/
type AlitripBtripFlightOrderSearchRequest struct {
    model.Params

    // 请求
    rq   *OpenSearchRq 

}

func NewAlitripBtripFlightOrderSearchRequest() *AlitripBtripFlightOrderSearchRequest{
    return &AlitripBtripFlightOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripFlightOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.flight.order.search"
}

func (r AlitripBtripFlightOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripFlightOrderSearchRequest) SetRq(rq *OpenSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripFlightOrderSearchRequest) GetRq() *OpenSearchRq {
    return r.rq
}

