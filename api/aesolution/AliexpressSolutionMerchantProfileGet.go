package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionmerchantprofileget aliexpress.solution.merchant.profile.get
// aliexpress.solution.merchant.profile.get
//
// API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
func Aliexpresssolutionmerchantprofileget(clt *core.SDKClient, req *aesolution.AliexpresssolutionmerchantprofilegetAPIRequest, session string) (*aesolution.AliexpresssolutionmerchantprofilegetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionmerchantprofilegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
