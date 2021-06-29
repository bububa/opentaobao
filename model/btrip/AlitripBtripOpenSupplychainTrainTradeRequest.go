package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅火车票交易流水接口 APIRequest
alitrip.btrip.open.supplychain.train.trade

商旅火车票交易流水接口
*/
type AlitripBtripOpenSupplychainTrainTradeRequest struct {
    model.Params

    // 入参
    rq   *OpenApiZzdSearchRq 

}

func NewAlitripBtripOpenSupplychainTrainTradeRequest() *AlitripBtripOpenSupplychainTrainTradeRequest{
    return &AlitripBtripOpenSupplychainTrainTradeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenSupplychainTrainTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.train.trade"
}

func (r AlitripBtripOpenSupplychainTrainTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenSupplychainTrainTradeRequest) SetRq(rq *OpenApiZzdSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenSupplychainTrainTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r.rq
}

