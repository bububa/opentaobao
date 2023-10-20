package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionbatchproductdelete aliexpress.solution.batch.product.delete
// aliexpress.solution.batch.product.delete
//
// Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
func Aliexpresssolutionbatchproductdelete(clt *core.SDKClient, req *aesolution.AliexpresssolutionbatchproductdeleteAPIRequest, session string) (*aesolution.AliexpresssolutionbatchproductdeleteAPIResponse, error) {
	var resp aesolution.AliexpresssolutionbatchproductdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
