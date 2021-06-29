package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部成本中心 API请求
alitrip.btip.cost.center.query

查询外部成本中心
*/
type AlitripBtipCostCenterQueryRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterQueryRq
}

// 初始化AlitripBtipCostCenterQueryRequest对象
func NewAlitripBtipCostCenterQueryRequest() *AlitripBtipCostCenterQueryRequest{
    return &AlitripBtipCostCenterQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtipCostCenterQueryRequest) GetApiMethodName() string {
    return "alitrip.btip.cost.center.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtipCostCenterQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtipCostCenterQueryRequest) SetRq(rq *OpenCostCenterQueryRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtipCostCenterQueryRequest) GetRq() *OpenCostCenterQueryRq {
    return r.rq
}
