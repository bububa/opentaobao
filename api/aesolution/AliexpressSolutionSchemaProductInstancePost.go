package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionSchemaProductInstancePost aliexpress.solution.schema.product.instance.post
// aliexpress.solution.schema.product.instance.post
//
// Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
func AliexpressSolutionSchemaProductInstancePost(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionSchemaProductInstancePostAPIRequest, resp *aesolution.AliexpressSolutionSchemaProductInstancePostAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
