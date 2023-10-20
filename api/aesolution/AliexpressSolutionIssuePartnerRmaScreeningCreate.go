package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaScreeningCreate aliexpress.solution.issue.partner.rma.screening.create
// aliexpress.solution.issue.partner.rma.screening.create
//
// Receives information about screening results from after sales partners
func AliexpressSolutionIssuePartnerRmaScreeningCreate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest, session string) (*aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse, error) {
	var resp aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
