package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加外部成本中心人员信息 APIRequest
alitrip.btrip.cost.center.entity.add

增加外部成本中心人员信息
*/
type AlitripBtripCostCenterEntityAddRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterAddEntityRq 

}

func NewAlitripBtripCostCenterEntityAddRequest() *AlitripBtripCostCenterEntityAddRequest{
    return &AlitripBtripCostCenterEntityAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterEntityAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.add"
}

func (r AlitripBtripCostCenterEntityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterEntityAddRequest) SetRq(rq *OpenCostCenterAddEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterEntityAddRequest) GetRq() *OpenCostCenterAddEntityRq {
    return r.rq
}

