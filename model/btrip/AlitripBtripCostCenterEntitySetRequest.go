package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.set

设置外部成本中心人员信息
*/
type AlitripBtripCostCenterEntitySetRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterSetEntityRq
}

// 初始化AlitripBtripCostCenterEntitySetRequest对象
func NewAlitripBtripCostCenterEntitySetRequest() *AlitripBtripCostCenterEntitySetRequest{
    return &AlitripBtripCostCenterEntitySetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntitySetRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.set"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntitySetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntitySetRequest) SetRq(rq *OpenCostCenterSetEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterEntitySetRequest) GetRq() *OpenCostCenterSetEntityRq {
    return r.rq
}
