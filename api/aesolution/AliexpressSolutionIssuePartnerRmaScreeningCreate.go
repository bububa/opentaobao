package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionissuepartnerrmascreeningcreate aliexpress.solution.issue.partner.rma.screening.create
// aliexpress.solution.issue.partner.rma.screening.create
//
// Receives information about screening results from after sales partners
func Aliexpresssolutionissuepartnerrmascreeningcreate(clt *core.SDKClient, req *aesolution.AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest, session string) (*aesolution.AliexpresssolutionissuepartnerrmascreeningcreateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionissuepartnerrmascreeningcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
