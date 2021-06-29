package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅火车票交易流水接口 API请求
alitrip.btrip.open.supplychain.train.trade

商旅火车票交易流水接口
*/
type AlitripBtripOpenSupplychainTrainTradeRequest struct {
    model.Params
    // 入参
    rq   *OpenApiZzdSearchRq
}

// 初始化AlitripBtripOpenSupplychainTrainTradeRequest对象
func NewAlitripBtripOpenSupplychainTrainTradeRequest() *AlitripBtripOpenSupplychainTrainTradeRequest{
    return &AlitripBtripOpenSupplychainTrainTradeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainTrainTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.train.trade"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainTrainTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainTrainTradeRequest) SetRq(rq *OpenApiZzdSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenSupplychainTrainTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r.rq
}
