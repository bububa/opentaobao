package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionproductschemaget get product schema
// aliexpress.solution.product.schema.get
//
// provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
func Aliexpresssolutionproductschemaget(clt *core.SDKClient, req *aesolution.AliexpresssolutionproductschemagetAPIRequest, session string) (*aesolution.AliexpresssolutionproductschemagetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionproductschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
