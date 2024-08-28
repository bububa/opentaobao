package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderInfoGet get order detail info
// aliexpress.solution.order.info.get
//
// get order detail info
func AliexpressSolutionOrderInfoGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderInfoGetAPIRequest, resp *aesolution.AliexpressSolutionOrderInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
