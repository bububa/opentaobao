package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改成本中心 API请求
alitrip.btrip.open.cost.center.modify

修改成本中心
*/
type AlitripBtripOpenCostCenterModifyRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterModifyRq
}

// 初始化AlitripBtripOpenCostCenterModifyRequest对象
func NewAlitripBtripOpenCostCenterModifyRequest() *AlitripBtripOpenCostCenterModifyRequest{
    return &AlitripBtripOpenCostCenterModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterModifyRequest) SetRq(_rq *OpenCostCenterModifyRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterModifyRequest) GetRq() *OpenCostCenterModifyRq {
    return r._rq
}
