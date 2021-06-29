package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索酒店订单列表 APIRequest
alitrip.btrip.hotel.order.search

企业获取商旅酒店订单数据
*/
type AlitripBtripHotelOrderSearchRequest struct {
    model.Params

    // 请求
    rq   *OpenSearchRq 

}

func NewAlitripBtripHotelOrderSearchRequest() *AlitripBtripHotelOrderSearchRequest{
    return &AlitripBtripHotelOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.order.search"
}

func (r AlitripBtripHotelOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelOrderSearchRequest) SetRq(rq *OpenSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripHotelOrderSearchRequest) GetRq() *OpenSearchRq {
    return r.rq
}

