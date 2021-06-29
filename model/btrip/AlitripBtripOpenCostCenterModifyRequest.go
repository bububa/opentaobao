package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改成本中心 APIRequest
alitrip.btrip.open.cost.center.modify

修改成本中心
*/
type AlitripBtripOpenCostCenterModifyRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterModifyRq 

}

func NewAlitripBtripOpenCostCenterModifyRequest() *AlitripBtripOpenCostCenterModifyRequest{
    return &AlitripBtripOpenCostCenterModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.modify"
}

func (r AlitripBtripOpenCostCenterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterModifyRequest) SetRq(rq *OpenCostCenterModifyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterModifyRequest) GetRq() *OpenCostCenterModifyRq {
    return r.rq
}

