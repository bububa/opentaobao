package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionproductpost Product posting API
// aliexpress.solution.product.post
//
// Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
func Aliexpresssolutionproductpost(clt *core.SDKClient, req *aesolution.AliexpresssolutionproductpostAPIRequest, session string) (*aesolution.AliexpresssolutionproductpostAPIResponse, error) {
	var resp aesolution.AliexpresssolutionproductpostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
