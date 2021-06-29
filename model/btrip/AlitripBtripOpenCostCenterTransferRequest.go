package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅成本中心转换为外部成本中心 APIRequest
alitrip.btrip.open.cost.center.transfer

商旅成本中心转换为外部成本中心
*/
type AlitripBtripOpenCostCenterTransferRequest struct {
    model.Params

    // 入参对象
    rq   *OpenCostCenterTransferRq 

}

func NewAlitripBtripOpenCostCenterTransferRequest() *AlitripBtripOpenCostCenterTransferRequest{
    return &AlitripBtripOpenCostCenterTransferRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenCostCenterTransferRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.cost.center.transfer"
}

func (r AlitripBtripOpenCostCenterTransferRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenCostCenterTransferRequest) SetRq(rq *OpenCostCenterTransferRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenCostCenterTransferRequest) GetRq() *OpenCostCenterTransferRq {
    return r.rq
}

