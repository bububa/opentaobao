package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionbatchproductpriceupdate aliexpress.solution.batch.product.price.update
// aliexpress.solution.batch.product.price.update
//
// batch product price update operation for oversea sellers
func Aliexpresssolutionbatchproductpriceupdate(clt *core.SDKClient, req *aesolution.AliexpresssolutionbatchproductpriceupdateAPIRequest, session string) (*aesolution.AliexpresssolutionbatchproductpriceupdateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionbatchproductpriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
