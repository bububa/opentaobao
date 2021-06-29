package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除成本中心人员信息 APIRequest
alitrip.btrip.open.cost.center.entity.delete

删除成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntityDeleteRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterDeleteEntityRq 

}

func NewAlitripBtripOpenCostCenterEntityDeleteRequest() *AlitripBtripOpenCostCenterEntityDeleteRequest{
    return &AlitripBtripOpenCostCenterEntityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterEntityDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.entity.delete"
}

func (r AlitripBtripOpenCostCenterEntityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterEntityDeleteRequest) SetRq(rq *OpenCostCenterDeleteEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterEntityDeleteRequest) GetRq() *OpenCostCenterDeleteEntityRq {
    return r.rq
}

