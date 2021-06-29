package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.add

增加外部成本中心人员信息
*/
type AlitripBtripCostCenterEntityAddRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterAddEntityRq
}

// 初始化AlitripBtripCostCenterEntityAddRequest对象
func NewAlitripBtripCostCenterEntityAddRequest() *AlitripBtripCostCenterEntityAddRequest{
    return &AlitripBtripCostCenterEntityAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntityAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.entity.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntityAddRequest) SetRq(rq *OpenCostCenterAddEntityRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterEntityAddRequest) GetRq() *OpenCostCenterAddEntityRq {
    return r.rq
}
