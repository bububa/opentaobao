package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdate aliexpress.solution.issue.partner.rma.reverselogistic.state.update
// aliexpress.solution.issue.partner.rma.reverselogistic.state.update
//
// Updates the reverse logistics state for after sales services
func AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
