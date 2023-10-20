package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaStateUpdate aliexpress.solution.issue.partner.rma.state.update
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
func AliexpressSolutionIssuePartnerRmaStateUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
