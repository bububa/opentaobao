package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAscpPricingScmTofAPIRequest
TOF&SCM营销域对接-成本录入设置 API请求
tmall.ascp.pricing.scm.tof

TOF&SCM营销域对接-成本录入设置 */
type TmallAscpPricingScmTofAPIRequest struct {
	model.Params
	// 成本价集合
	_costs []ItemSkuCost
}

// New
