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
type AlitripBtripCostCenterDeleteRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterDeleteRq
}

// 初始化AlitripBtripCostCenterDeleteRequest对象
func NewAlitripBtripCostCenterDeleteRequest() *AlitripBtripCostCenterDeleteRequest{
    return &AlitripBtripCostCenterDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterDeleteRequest) SetRq(rq *OpenCostCenterDeleteRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterDeleteRequest) GetRq() *OpenCostCenterDeleteRq {
    return r.rq
}
