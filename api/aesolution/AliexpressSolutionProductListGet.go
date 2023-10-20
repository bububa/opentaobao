package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionproductlistget Get product list
// aliexpress.solution.product.list.get
//
// Get product list
func Aliexpresssolutionproductlistget(clt *core.SDKClient, req *aesolution.AliexpresssolutionproductlistgetAPIRequest, session string) (*aesolution.AliexpresssolutionproductlistgetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionproductlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
