package ascp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TOF&SCM营销域对接-成本录入设置 API请求
tmall.ascp.pricing.scm.tof

TOF&SCM营销域对接-成本录入设置
*/
type TmallAscpPricingScmTofRequest struct {
    model.Params
    // 成本价集合
    _costs   []ItemSkuCost
}

// 初始化TmallAscpPricingScmTofRequest对象
func NewTmallAscpPricingScmTofRequest() *TmallAscpPricingScmTofRequest{
    return &TmallAscpPricingScmTofRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAscpPricingScmTofRequest) GetApiMethodName() string {
    return "tmall.ascp.pricing.scm.tof"
}

// IRequest interface 方法, 获取API参数
func (r TmallAscpPricingScmTofRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Costs Setter
// 成本价集合
func (r *TmallAscpPricingScmTofRequest) SetCosts(_costs []ItemSkuCost) error {
    r._costs = _costs
    r.Set("costs", _costs)
    return nil
}

// Costs Getter
func (r TmallAscpPricingScmTofRequest) GetCosts() []ItemSkuCost {
    return r._costs
}
