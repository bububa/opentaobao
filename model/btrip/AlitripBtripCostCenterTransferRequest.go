package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅成本中心转换为外部成本中心 APIRequest
alitrip.btrip.cost.center.transfer

商旅成本中心转换为外部成本中心
*/
type AlitripBtripCostCenterTransferRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterTransferRq 

}

func NewAlitripBtripCostCenterTransferRequest() *AlitripBtripCostCenterTransferRequest{
    return &AlitripBtripCostCenterTransferRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterTransferRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.transfer"
}

func (r AlitripBtripCostCenterTransferRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterTransferRequest) SetRq(rq *OpenCostCenterTransferRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterTransferRequest) GetRq() *OpenCostCenterTransferRq {
    return r.rq
}

