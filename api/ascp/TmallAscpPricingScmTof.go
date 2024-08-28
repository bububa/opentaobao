package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TmallAscpPricingScmTof TOF&SCM营销域对接-成本录入设置
// tmall.ascp.pricing.scm.tof
//
// TOF&amp;SCM营销域对接-成本录入设置
func TmallAscpPricingScmTof(ctx context.Context, clt *core.SDKClient, req *ascp.TmallAscpPricingScmTofAPIRequest, resp *ascp.TmallAscpPricingScmTofAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
