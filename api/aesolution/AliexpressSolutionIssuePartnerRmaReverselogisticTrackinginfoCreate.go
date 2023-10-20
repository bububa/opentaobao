package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreate aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
// aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
//
// Receives information about reverse logistics tracking info
func AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
