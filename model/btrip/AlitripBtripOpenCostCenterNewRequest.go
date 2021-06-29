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
type AlitripBtripOpenCostCenterNewRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterSaveRq
}

// 初始化AlitripBtripOpenCostCenterNewRequest对象
func NewAlitripBtripOpenCostCenterNewRequest() *AlitripBtripOpenCostCenterNewRequest{
    return &AlitripBtripOpenCostCenterNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterNewRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.new"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterNewRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterNewRequest) GetRq() *OpenCostCenterSaveRq {
    return r._rq
}
