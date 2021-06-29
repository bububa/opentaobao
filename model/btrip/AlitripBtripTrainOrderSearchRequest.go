package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取火车票订单列表 APIRequest
alitrip.btrip.train.order.search

第三方OA厂商获取自己的火车票数据
*/
type AlitripBtripTrainOrderSearchRequest struct {
    model.Params

    // 请求
    rq   *OpenSearchRq 

}

func NewAlitripBtripTrainOrderSearchRequest() *AlitripBtripTrainOrderSearchRequest{
    return &AlitripBtripTrainOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripTrainOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.train.order.search"
}

func (r AlitripBtripTrainOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripTrainOrderSearchRequest) SetRq(rq *OpenSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripTrainOrderSearchRequest) GetRq() *OpenSearchRq {
    return r.rq
}

