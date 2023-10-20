package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAscpPricingScmTofAPIRequest TOF&SCM营销域对接-成本录入设置 API请求
// tmall.ascp.pricing.scm.tof
//
// TOF&amp;SCM营销域对接-成本录入设置
type TmallAscpPricingScmTofAPIRequest struct {
	model.Params
	// 成本价集合
	_costs []ItemSkuCost
}

// NewTmallAscpPricingScmTofRequest 初始化TmallAscpPricingScmTofAPIRequest对象
func NewTmallAscpPricingScmTofRequest() *TmallAscpPricingScmTofAPIRequest {
	return &TmallAscpPricingScmTofAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAscpPricingScmTofAPIRequest) Reset() {
	r._costs = r._costs[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAscpPricingScmTofAPIRequest) GetApiMethodName() string {
	return "tmall.ascp.pricing.scm.tof"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAscpPricingScmTofAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAscpPricingScmTofAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCosts is Costs Setter
// 成本价集合
func (r *TmallAscpPricingScmTofAPIRequest) SetCosts(_costs []ItemSkuCost) error {
	r._costs = _costs
	r.Set("costs", _costs)
	return nil
}

// GetCosts Costs Getter
func (r TmallAscpPricingScmTofAPIRequest) GetCosts() []ItemSkuCost {
	return r._costs
}

var poolTmallAscpPricingScmTofAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAscpPricingScmTofRequest()
	},
}

// GetTmallAscpPricingScmTofRequest 从 sync.Pool 获取 TmallAscpPricingScmTofAPIRequest
func GetTmallAscpPricingScmTofAPIRequest() *TmallAscpPricingScmTofAPIRequest {
	return poolTmallAscpPricingScmTofAPIRequest.Get().(*TmallAscpPricingScmTofAPIRequest)
}

// ReleaseTmallAscpPricingScmTofAPIRequest 将 TmallAscpPricingScmTofAPIRequest 放入 sync.Pool
func ReleaseTmallAscpPricingScmTofAPIRequest(v *TmallAscpPricingScmTofAPIRequest) {
	v.Reset()
	poolTmallAscpPricingScmTofAPIRequest.Put(v)
}
