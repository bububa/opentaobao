package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置成本中心人员信息 API请求
alitrip.btrip.open.cost.center.entity.set

设置成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntitySetRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterSetEntityRq
}

// 初始化AlitripBtripOpenCostCenterEntitySetRequest对象
func NewAlitripBtripOpenCostCenterEntitySetRequest() *AlitripBtripOpenCostCenterEntitySetRequest{
    return &AlitripBtripOpenCostCenterEntitySetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntitySetRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.entity.set"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntitySetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntitySetRequest) SetRq(_rq *OpenCostCenterSetEntityRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterEntitySetRequest) GetRq() *OpenCostCenterSetEntityRq {
    return r._rq
}
