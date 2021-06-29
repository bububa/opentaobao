package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅成本中心转换为外部成本中心 API请求
alitrip.btrip.open.cost.center.transfer

商旅成本中心转换为外部成本中心
*/
type AlitripBtripOpenCostCenterTransferRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenCostCenterTransferRq
}

// 初始化AlitripBtripOpenCostCenterTransferRequest对象
func NewAlitripBtripOpenCostCenterTransferRequest() *AlitripBtripOpenCostCenterTransferRequest{
    return &AlitripBtripOpenCostCenterTransferRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterTransferRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.transfer"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterTransferRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterTransferRequest) SetRq(_rq *OpenCostCenterTransferRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenCostCenterTransferRequest) GetRq() *OpenCostCenterTransferRq {
    return r._rq
}
