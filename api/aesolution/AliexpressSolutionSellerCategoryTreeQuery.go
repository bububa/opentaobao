package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionSellerCategoryTreeQuery aliexpress.solution.seller.category.tree.query
// aliexpress.solution.seller.category.tree.query
//
// API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
func AliexpressSolutionSellerCategoryTreeQuery(clt *core.SDKClient, req *aesolution.AliexpressSolutionSellerCategoryTreeQueryAPIRequest, resp *aesolution.AliexpressSolutionSellerCategoryTreeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
