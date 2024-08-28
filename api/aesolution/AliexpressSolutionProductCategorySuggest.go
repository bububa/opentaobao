package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductCategorySuggest Suggest product categories
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
func AliexpressSolutionProductCategorySuggest(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionProductCategorySuggestAPIRequest, resp *aesolution.AliexpressSolutionProductCategorySuggestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
