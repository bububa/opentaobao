package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallascppricingscmtofAPIRequest TOF&SCM营销域对接-成本录入设置 API请求
// tmall.ascp.pricing.scm.tof
//
// TOF&amp;SCM营销域对接-成本录入设置
type TmallascppricingscmtofAPIRequest struct {
	model.Params
	// 成本价集合
	_costs []ItemSkuCost
}

// NewTmallascppricingscmtofRequest 初始化TmallascppricingscmtofAPIRequest对象
func NewTmallascppricingscmtofRequest() *TmallascppricingscmtofAPIRequest {
	return &TmallascppricingscmtofAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallascppricingscmtofAPIRequest) GetApiMethodName() string {
	return "tmall.ascp.pricing.scm.tof"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallascppricingscmtofAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallascppricingscmtofAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCosts is Costs Setter
// 成本价集合
func (r *TmallascppricingscmtofAPIRequest) SetCosts(_costs []ItemSkuCost) error {
	r._costs = _costs
	r.Set("costs", _costs)
	return nil
}

// GetCosts Costs Getter
func (r TmallascppricingscmtofAPIRequest) GetCosts() []ItemSkuCost {
	return r._costs
}
