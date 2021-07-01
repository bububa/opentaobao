package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest
aliexpress.solution.issue.partner.rma.state.update API请求
aliexpress.solution.issue.partner.rma.state.update

Receive changes in state updates for RMAs orders from after sales partners */
type AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest struct {
	model.Params
	// RMA's order state update request
	_rmaStateUpdateRequest *RmaStateUpdateRequest
}

// New
