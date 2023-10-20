package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionSkuAttributeQuery Query the sku attribute information belonged to a specific category
// aliexpress.solution.sku.attribute.query
//
// Query the sku attribute information belonged to a specific category, customized for oversea merchants.
func AliexpressSolutionSkuAttributeQuery(clt *core.SDKClient, req *aesolution.AliexpressSolutionSkuAttributeQueryAPIRequest, resp *aesolution.AliexpressSolutionSkuAttributeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
