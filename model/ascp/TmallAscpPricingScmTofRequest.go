package ascp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TOF&SCM营销域对接-成本录入设置 APIRequest
tmall.ascp.pricing.scm.tof

TOF&SCM营销域对接-成本录入设置
*/
type TmallAscpPricingScmTofRequest struct {
    model.Params

    // 成本价集合
    costs   []ItemSkuCost 

}

func NewTmallAscpPricingScmTofRequest() *TmallAscpPricingScmTofRequest{
    return &TmallAscpPricingScmTofRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAscpPricingScmTofRequest) GetApiMethodName() string {
    return "tmall.ascp.pricing.scm.tof"
}

func (r TmallAscpPricingScmTofRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAscpPricingScmTofRequest) SetCosts(costs []ItemSkuCost) error {
    r.costs = costs
    r.Set("costs", costs)
    return nil
}

func (r TmallAscpPricingScmTofRequest) GetCosts() []ItemSkuCost {
    return r.costs
}

