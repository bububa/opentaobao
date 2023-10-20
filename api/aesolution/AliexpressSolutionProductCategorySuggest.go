package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductCategorySuggest Suggest product categories
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
func AliexpressSolutionProductCategorySuggest(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductCategorySuggestAPIRequest, session string) (*aesolution.AliexpressSolutionProductCategorySuggestAPIResponse, error) {
	var resp aesolution.AliexpressSolutionProductCategorySuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
