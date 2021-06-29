package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下pr单 APIRequest
alibaba.pur.pr.create

下pr单
*/
type AlibabaPurPrCreateRequest struct {
    model.Params

    // 订单信息
    purReq   *MallReceivePrRequest 

}

func NewAlibabaPurPrCreateRequest() *AlibabaPurPrCreateRequest{
    return &AlibabaPurPrCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurPrCreateRequest) GetApiMethodName() string {
    return "alibaba.pur.pr.create"
}

func (r AlibabaPurPrCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurPrCreateRequest) SetPurReq(purReq *MallReceivePrRequest) error {
    r.purReq = purReq
    r.Set("pur_req", purReq)
    return nil
}

func (r AlibabaPurPrCreateRequest) GetPurReq() *MallReceivePrRequest {
    return r.purReq
}

