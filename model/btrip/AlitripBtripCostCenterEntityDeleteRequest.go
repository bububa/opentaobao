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
type AlitripBtripCostCenterEntityDeleteRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterDeleteEntityRq
}

// 初始化AlitripBtripCostCenterEntityDeleteRequest对象
func NewAlitripBtripCostCenterEntityDeleteRequest() *AlitripBtripCostCenterEntityDeleteRequest{
    return &AlitripBtripCostCenterEntityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntityDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntityDeleteRequest) SetRq(rq *OpenCostCenterDeleteEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterEntityDeleteRequest) GetRq() *OpenCostCenterDeleteEntityRq {
    return r.rq
}
