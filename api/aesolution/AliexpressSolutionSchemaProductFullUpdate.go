package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionSchemaProductFullUpdate aliexpress.solution.schema.product.full.update
// aliexpress.solution.schema.product.full.update
//
// Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
func AliexpressSolutionSchemaProductFullUpdate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionSchemaProductFullUpdateAPIRequest, resp *aesolution.AliexpressSolutionSchemaProductFullUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
