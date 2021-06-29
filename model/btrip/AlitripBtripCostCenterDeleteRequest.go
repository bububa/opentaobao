package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除外部成本中心 APIRequest
alitrip.btrip.cost.center.delete

删除外部成本中心
*/
type AlitripBtripCostCenterDeleteRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterDeleteRq 

}

func NewAlitripBtripCostCenterDeleteRequest() *AlitripBtripCostCenterDeleteRequest{
    return &AlitripBtripCostCenterDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.delete"
}

func (r AlitripBtripCostCenterDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterDeleteRequest) SetRq(rq *OpenCostCenterDeleteRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterDeleteRequest) GetRq() *OpenCostCenterDeleteRq {
    return r.rq
}

