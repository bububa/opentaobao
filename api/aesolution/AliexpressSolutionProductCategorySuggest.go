package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductCategorySuggest Suggest product categories
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
func AliexpressSolutionProductCategorySuggest(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductCategorySuggestAPIRequest, resp *aesolution.AliexpressSolutionProductCategorySuggestAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
