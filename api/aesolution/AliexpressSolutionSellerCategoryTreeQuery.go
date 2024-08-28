package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionSellerCategoryTreeQuery aliexpress.solution.seller.category.tree.query
// aliexpress.solution.seller.category.tree.query
//
// API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
func AliexpressSolutionSellerCategoryTreeQuery(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionSellerCategoryTreeQueryAPIRequest, resp *aesolution.AliexpressSolutionSellerCategoryTreeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
