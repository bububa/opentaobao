package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改外部成本中心 APIRequest
alitrip.btrip.cost.center.modify

修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
*/
type AlitripBtripCostCenterModifyRequest struct {
    model.Params

    // 请求对象
    rq   *OpenCostCenterModifyRq 

}

func NewAlitripBtripCostCenterModifyRequest() *AlitripBtripCostCenterModifyRequest{
    return &AlitripBtripCostCenterModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.modify"
}

func (r AlitripBtripCostCenterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterModifyRequest) SetRq(rq *OpenCostCenterModifyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCostCenterModifyRequest) GetRq() *OpenCostCenterModifyRq {
    return r.rq
}

