package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionskuattributequery Query the sku attribute information belonged to a specific category
// aliexpress.solution.sku.attribute.query
//
// Query the sku attribute information belonged to a specific category, customized for oversea merchants.
func Aliexpresssolutionskuattributequery(clt *core.SDKClient, req *aesolution.AliexpresssolutionskuattributequeryAPIRequest, session string) (*aesolution.AliexpresssolutionskuattributequeryAPIResponse, error) {
	var resp aesolution.AliexpresssolutionskuattributequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
