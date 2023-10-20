package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductEdit Edit Product API
// aliexpress.solution.product.edit
//
// API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
func AliexpressSolutionProductEdit(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductEditAPIRequest, resp *aesolution.AliexpressSolutionProductEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
