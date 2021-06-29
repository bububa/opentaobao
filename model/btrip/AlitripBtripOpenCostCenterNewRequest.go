package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增成本中心 APIRequest
alitrip.btrip.open.cost.center.new

新增成本中心
*/
type AlitripBtripOpenCostCenterNewRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterSaveRq 

}

func NewAlitripBtripOpenCostCenterNewRequest() *AlitripBtripOpenCostCenterNewRequest{
    return &AlitripBtripOpenCostCenterNewRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterNewRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.new"
}

func (r AlitripBtripOpenCostCenterNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterNewRequest) SetRq(rq *OpenCostCenterSaveRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterNewRequest) GetRq() *OpenCostCenterSaveRq {
    return r.rq
}

