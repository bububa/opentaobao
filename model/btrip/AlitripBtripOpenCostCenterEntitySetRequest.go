package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置成本中心人员信息 APIRequest
alitrip.btrip.open.cost.center.entity.set

设置成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntitySetRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterSetEntityRq 

}

func NewAlitripBtripOpenCostCenterEntitySetRequest() *AlitripBtripOpenCostCenterEntitySetRequest{
    return &AlitripBtripOpenCostCenterEntitySetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterEntitySetRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.entity.set"
}

func (r AlitripBtripOpenCostCenterEntitySetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterEntitySetRequest) SetRq(rq *OpenCostCenterSetEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterEntitySetRequest) GetRq() *OpenCostCenterSetEntityRq {
    return r.rq
}

