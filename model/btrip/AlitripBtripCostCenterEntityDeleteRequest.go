package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.delete

删除外部成本中心人员信息
*/
type AlitripBtripCostCenterEntityDeleteAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenCostCenterDeleteEntityRq
}

// 初始化AlitripBtripCostCenterEntityDeleteAPIRequest对象
func NewAlitripBtripCostCenterEntityDeleteRequest() *AlitripBtripCostCenterEntityDeleteAPIRequest{
    return &AlitripBtripCostCenterEntityDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntityDeleteAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntityDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntityDeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteEntityRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterEntityDeleteAPIRequest) GetRq() *OpenCostCenterDeleteEntityRq {
    return r._rq
}
