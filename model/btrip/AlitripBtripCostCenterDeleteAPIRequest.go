package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除外部成本中心 API请求
alitrip.btrip.cost.center.delete

删除外部成本中心
*/
type AlitripBtripCostCenterDeleteAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenCostCenterDeleteRq
}

// 初始化AlitripBtripCostCenterDeleteAPIRequest对象
func NewAlitripBtripCostCenterDeleteRequest() *AlitripBtripCostCenterDeleteAPIRequest{
    return &AlitripBtripCostCenterDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterDeleteAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterDeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterDeleteAPIRequest) GetRq() *OpenCostCenterDeleteRq {
    return r._rq
}
