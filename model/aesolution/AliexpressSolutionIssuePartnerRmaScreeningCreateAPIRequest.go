package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest
aliexpress.solution.issue.partner.rma.screening.create API请求
aliexpress.solution.issue.partner.rma.screening.create

Receives information about screening results from after sales partners */
type AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest struct {
	model.Params
	// Screening result creation request
	_screeningResultCreationRequest *RmaScreeningCreationRequest
}

// New
