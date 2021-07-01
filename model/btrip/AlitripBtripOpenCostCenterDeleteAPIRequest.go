package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除成本中心 API请求
alitrip.btrip.open.cost.center.delete

删除成本中心
*/
type AlitripBtripOpenCostCenterDeleteAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterDeleteRq
}

// 初始化AlitripBtripOpenCostCenterDeleteAPIRequest对象
func NewAlitripBtripOpenCostCenterDeleteRequest() *AlitripBtripOpenCostCenterDeleteAPIRequest{
    return &AlitripBtripOpenCostCenterDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterDeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetRq() *OpenCostCenterDeleteRq {
    return r._rq
}
