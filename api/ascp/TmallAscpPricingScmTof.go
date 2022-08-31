package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TmallAscpPricingScmTof TOF&SCM营销域对接-成本录入设置
// tmall.ascp.pricing.scm.tof
//
// TOF&amp;SCM营销域对接-成本录入设置
func TmallAscpPricingScmTof(clt *core.SDKClient, req *ascp.TmallAscpPricingScmTofAPIRequest, session string) (*ascp.TmallAscpPricingScmTofAPIResponse, error) {
	var resp ascp.TmallAscpPricingScmTofAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
