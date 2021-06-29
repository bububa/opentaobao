package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建外部成本中心 APIRequest
alitrip.btrip.cost.center.new

新建外部成本中心
*/
type AlitripBtripCostCenterNewRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterSaveRq 

}

func NewAlitripBtripCostCenterNewRequest() *AlitripBtripCostCenterNewRequest{
    return &AlitripBtripCostCenterNewRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterNewRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.new"
}

func (r AlitripBtripCostCenterNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterNewRequest) SetRq(rq *OpenCostCenterSaveRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterNewRequest) GetRq() *OpenCostCenterSaveRq {
    return r.rq
}

