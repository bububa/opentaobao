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
type TmallAscpPricingScmTofAPIRequest struct {
    model.Params
    // 成本价集合
    _costs   []ItemSkuCost
}

// 初始化TmallAscpPricingScmTofAPIRequest对象
func NewTmallAscpPricingScmTofRequest() *TmallAscpPricingScmTofAPIRequest{
    return &TmallAscpPricingScmTofAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAscpPricingScmTofAPIRequest) GetApiMethodName() string {
    return "tmall.ascp.pricing.scm.tof"
}

// IRequest interface 方法, 获取API参数
func (r TmallAscpPricingScmTofAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Costs Setter
// 成本价集合
func (r *TmallAscpPricingScmTofAPIRequest) SetCosts(_costs []ItemSkuCost) error {
    r._costs = _costs
    r.Set("costs", _costs)
    return nil
}

// Costs Getter
func (r TmallAscpPricingScmTofAPIRequest) GetCosts() []ItemSkuCost {
    return r._costs
}
