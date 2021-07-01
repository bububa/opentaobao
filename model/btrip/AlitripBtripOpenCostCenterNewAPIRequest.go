package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增成本中心 API请求
alitrip.btrip.open.cost.center.new

新增成本中心
*/
type AlitripBtripOpenCostCenterNewAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterSaveRq
}

// 初始化AlitripBtripOpenCostCenterNewAPIRequest对象
func NewAlitripBtripOpenCostCenterNewRequest() *AlitripBtripOpenCostCenterNewAPIRequest{
    return &AlitripBtripOpenCostCenterNewAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.new"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterNewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetRq() *OpenCostCenterSaveRq {
    return r._rq
}
