package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create API请求
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create

Receives information about reverse logistics tracking info */
type AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest struct {
	model.Params
	// Logistic's order creation request
	_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest
}

// New
