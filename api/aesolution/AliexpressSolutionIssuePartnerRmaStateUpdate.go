package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaStateUpdate aliexpress.solution.issue.partner.rma.state.update
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
func AliexpressSolutionIssuePartnerRmaStateUpdate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
