package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionorderfulfill fulfill order
// aliexpress.solution.order.fulfill
//
// fulfill order for seller
func Aliexpresssolutionorderfulfill(clt *core.SDKClient, req *aesolution.AliexpresssolutionorderfulfillAPIRequest, session string) (*aesolution.AliexpresssolutionorderfulfillAPIResponse, error) {
	var resp aesolution.AliexpresssolutionorderfulfillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
