package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionissuepartnerrmareverselogisticstateupdate aliexpress.solution.issue.partner.rma.reverselogistic.state.update
// aliexpress.solution.issue.partner.rma.reverselogistic.state.update
//
// Updates the reverse logistics state for after sales services
func Aliexpresssolutionissuepartnerrmareverselogisticstateupdate(clt *core.SDKClient, req *aesolution.AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest, session string) (*aesolution.AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
