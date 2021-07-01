package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest
aliexpress.solution.issue.partner.rma.reverselogistic.state.update API请求
aliexpress.solution.issue.partner.rma.reverselogistic.state.update

Updates the reverse logistics state for after sales services */
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest struct {
	model.Params
	// Logistic order state update request
	_logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest
}

// New
