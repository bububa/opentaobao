package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询成本中心 API请求
alitrip.btrip.open.cost.center.query

查询成本中心
*/
type AlitripBtripOpenCostCenterQueryRequest struct {
    model.Params
    // 入参对象
    rq   *OpenCostCenterQueryRq
}

// 初始化AlitripBtripOpenCostCenterQueryRequest对象
func NewAlitripBtripOpenCostCenterQueryRequest() *AlitripBtripOpenCostCenterQueryRequest{
    return &AlitripBtripOpenCostCenterQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterQueryRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterQueryRequest) SetRq(rq *OpenCostCenterQueryRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterQueryRequest) GetRq() *OpenCostCenterQueryRq {
    return r.rq
}
