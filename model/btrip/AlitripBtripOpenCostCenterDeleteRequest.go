package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除成本中心 APIRequest
alitrip.btrip.open.cost.center.delete

删除成本中心
*/
type AlitripBtripOpenCostCenterDeleteRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterDeleteRq 

}

func NewAlitripBtripOpenCostCenterDeleteRequest() *AlitripBtripOpenCostCenterDeleteRequest{
    return &AlitripBtripOpenCostCenterDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.delete"
}

func (r AlitripBtripOpenCostCenterDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterDeleteRequest) SetRq(rq *OpenCostCenterDeleteRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterDeleteRequest) GetRq() *OpenCostCenterDeleteRq {
    return r.rq
}

