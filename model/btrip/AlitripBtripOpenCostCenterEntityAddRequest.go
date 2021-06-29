package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加成本中心人员信息 API请求
alitrip.btrip.open.cost.center.entity.add

增加成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntityAddRequest struct {
    model.Params
    // 入参对象
    rq   *OpenCostCenterAddEntityRq
}

// 初始化AlitripBtripOpenCostCenterEntityAddRequest对象
func NewAlitripBtripOpenCostCenterEntityAddRequest() *AlitripBtripOpenCostCenterEntityAddRequest{
    return &AlitripBtripOpenCostCenterEntityAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntityAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.entity.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntityAddRequest) SetRq(rq *OpenCostCenterAddEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterEntityAddRequest) GetRq() *OpenCostCenterAddEntityRq {
    return r.rq
}
