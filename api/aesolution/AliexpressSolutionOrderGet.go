package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionorderget get order list
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
func Aliexpresssolutionorderget(clt *core.SDKClient, req *aesolution.AliexpresssolutionordergetAPIRequest, session string) (*aesolution.AliexpresssolutionordergetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
