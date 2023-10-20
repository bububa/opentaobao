package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionproductedit Edit Product API
// aliexpress.solution.product.edit
//
// API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
func Aliexpresssolutionproductedit(clt *core.SDKClient, req *aesolution.AliexpresssolutionproducteditAPIRequest, session string) (*aesolution.AliexpresssolutionproducteditAPIResponse, error) {
	var resp aesolution.AliexpresssolutionproducteditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
