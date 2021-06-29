package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加成本中心人员信息 APIRequest
alitrip.btrip.open.cost.center.entity.add

增加成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntityAddRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterAddEntityRq 

}

func NewAlitripBtripOpenCostCenterEntityAddRequest() *AlitripBtripOpenCostCenterEntityAddRequest{
    return &AlitripBtripOpenCostCenterEntityAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterEntityAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.entity.add"
}

func (r AlitripBtripOpenCostCenterEntityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterEntityAddRequest) SetRq(rq *OpenCostCenterAddEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterEntityAddRequest) GetRq() *OpenCostCenterAddEntityRq {
    return r.rq
}

