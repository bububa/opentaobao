package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建外部成本中心 API请求
alitrip.btrip.cost.center.new

新建外部成本中心
*/
type AlitripBtripCostCenterNewRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterSaveRq
}

// 初始化AlitripBtripCostCenterNewRequest对象
func NewAlitripBtripCostCenterNewRequest() *AlitripBtripCostCenterNewRequest{
    return &AlitripBtripCostCenterNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterNewRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.new"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterNewRequest) SetRq(rq *OpenCostCenterSaveRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterNewRequest) GetRq() *OpenCostCenterSaveRq {
    return r.rq
}
