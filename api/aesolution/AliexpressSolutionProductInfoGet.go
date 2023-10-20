package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionproductinfoget Get Single Product Info
// aliexpress.solution.product.info.get
//
// Get Single Product Info
func Aliexpresssolutionproductinfoget(clt *core.SDKClient, req *aesolution.AliexpresssolutionproductinfogetAPIRequest, session string) (*aesolution.AliexpresssolutionproductinfogetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionproductinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
