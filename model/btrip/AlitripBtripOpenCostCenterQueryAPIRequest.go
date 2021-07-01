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
type AlitripBtripOpenCostCenterQueryAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterQueryRq
}

// 初始化AlitripBtripOpenCostCenterQueryAPIRequest对象
func NewAlitripBtripOpenCostCenterQueryRequest() *AlitripBtripOpenCostCenterQueryAPIRequest{
    return &AlitripBtripOpenCostCenterQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterQueryAPIRequest) SetRq(_rq *OpenCostCenterQueryRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetRq() *OpenCostCenterQueryRq {
    return r._rq
}
