package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除成本中心人员信息 API请求
alitrip.btrip.open.cost.center.entity.delete

删除成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntityDeleteAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterDeleteEntityRq
}

// 初始化AlitripBtripOpenCostCenterEntityDeleteAPIRequest对象
func NewAlitripBtripOpenCostCenterEntityDeleteRequest() *AlitripBtripOpenCostCenterEntityDeleteAPIRequest{
    return &AlitripBtripOpenCostCenterEntityDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.entity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntityDeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteEntityRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetRq() *OpenCostCenterDeleteEntityRq {
    return r._rq
}
