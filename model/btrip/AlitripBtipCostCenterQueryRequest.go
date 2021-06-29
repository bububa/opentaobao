package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部成本中心 APIRequest
alitrip.btip.cost.center.query

查询外部成本中心
*/
type AlitripBtipCostCenterQueryRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterQueryRq 

}

func NewAlitripBtipCostCenterQueryRequest() *AlitripBtipCostCenterQueryRequest{
    return &AlitripBtipCostCenterQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtipCostCenterQueryRequest) GetApiMethodName() string {
    return "alitrip.btip.cost.center.query"
}

func (r AlitripBtipCostCenterQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtipCostCenterQueryRequest) SetRq(rq *OpenCostCenterQueryRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtipCostCenterQueryRequest) GetRq() *OpenCostCenterQueryRq {
    return r.rq
}

