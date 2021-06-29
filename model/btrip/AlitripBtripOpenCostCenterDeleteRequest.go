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
type AlitripBtripOpenCostCenterDeleteRequest struct {
    model.Params
    // 入参对象
    rq   *OpenCostCenterDeleteRq
}

// 初始化AlitripBtripOpenCostCenterDeleteRequest对象
func NewAlitripBtripOpenCostCenterDeleteRequest() *AlitripBtripOpenCostCenterDeleteRequest{
    return &AlitripBtripOpenCostCenterDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterDeleteRequest) SetRq(rq *OpenCostCenterDeleteRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterDeleteRequest) GetRq() *OpenCostCenterDeleteRq {
    return r.rq
}
