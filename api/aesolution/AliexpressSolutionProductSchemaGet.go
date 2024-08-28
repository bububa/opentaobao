package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductSchemaGet get product schema
// aliexpress.solution.product.schema.get
//
// provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
func AliexpressSolutionProductSchemaGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionProductSchemaGetAPIRequest, resp *aesolution.AliexpressSolutionProductSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
