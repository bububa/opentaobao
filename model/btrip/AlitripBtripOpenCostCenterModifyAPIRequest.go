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
type AlitripBtripOpenCostCenterModifyAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterModifyRq
}

// 初始化AlitripBtripOpenCostCenterModifyAPIRequest对象
func NewAlitripBtripOpenCostCenterModifyRequest() *AlitripBtripOpenCostCenterModifyAPIRequest{
    return &AlitripBtripOpenCostCenterModifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterModifyAPIRequest) SetRq(_rq *OpenCostCenterModifyRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetRq() *OpenCostCenterModifyRq {
    return r._rq
}
