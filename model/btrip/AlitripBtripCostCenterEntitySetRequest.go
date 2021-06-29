package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置外部成本中心人员信息 APIRequest
alitrip.btrip.cost.center.entity.set

设置外部成本中心人员信息
*/
type AlitripBtripCostCenterEntitySetRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterSetEntityRq 

}

func NewAlitripBtripCostCenterEntitySetRequest() *AlitripBtripCostCenterEntitySetRequest{
    return &AlitripBtripCostCenterEntitySetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterEntitySetRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.set"
}

func (r AlitripBtripCostCenterEntitySetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterEntitySetRequest) SetRq(rq *OpenCostCenterSetEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterEntitySetRequest) GetRq() *OpenCostCenterSetEntityRq {
    return r.rq
}

