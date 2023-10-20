package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaScreeningCreate aliexpress.solution.issue.partner.rma.screening.create
// aliexpress.solution.issue.partner.rma.screening.create
//
// Receives information about screening results from after sales partners
func AliexpressSolutionIssuePartnerRmaScreeningCreate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
