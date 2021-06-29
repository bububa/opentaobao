package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询成本中心 APIRequest
alitrip.btrip.open.cost.center.query

查询成本中心
*/
type AlitripBtripOpenCostCenterQueryRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterQueryRq 

}

func NewAlitripBtripOpenCostCenterQueryRequest() *AlitripBtripOpenCostCenterQueryRequest{
    return &AlitripBtripOpenCostCenterQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterQueryRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.query"
}

func (r AlitripBtripOpenCostCenterQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterQueryRequest) SetRq(rq *OpenCostCenterQueryRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterQueryRequest) GetRq() *OpenCostCenterQueryRq {
    return r.rq
}

