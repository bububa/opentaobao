package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionissuepartnerrmastateupdate aliexpress.solution.issue.partner.rma.state.update
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
func Aliexpresssolutionissuepartnerrmastateupdate(clt *core.SDKClient, req *aesolution.AliexpresssolutionissuepartnerrmastateupdateAPIRequest, session string) (*aesolution.AliexpresssolutionissuepartnerrmastateupdateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionissuepartnerrmastateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
