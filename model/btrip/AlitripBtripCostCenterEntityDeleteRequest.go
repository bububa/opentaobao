package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除外部成本中心人员信息 APIRequest
alitrip.btrip.cost.center.entity.delete

删除外部成本中心人员信息
*/
type AlitripBtripCostCenterEntityDeleteRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterDeleteEntityRq 

}

func NewAlitripBtripCostCenterEntityDeleteRequest() *AlitripBtripCostCenterEntityDeleteRequest{
    return &AlitripBtripCostCenterEntityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterEntityDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.delete"
}

func (r AlitripBtripCostCenterEntityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterEntityDeleteRequest) SetRq(rq *OpenCostCenterDeleteEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterEntityDeleteRequest) GetRq() *OpenCostCenterDeleteEntityRq {
    return r.rq
}

