package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionproductcategorysuggest Suggest product categories
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
func Aliexpresssolutionproductcategorysuggest(clt *core.SDKClient, req *aesolution.AliexpresssolutionproductcategorysuggestAPIRequest, session string) (*aesolution.AliexpresssolutionproductcategorysuggestAPIResponse, error) {
	var resp aesolution.AliexpresssolutionproductcategorysuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
