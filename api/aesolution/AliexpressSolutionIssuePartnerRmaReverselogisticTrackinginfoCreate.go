package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

/* AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreate
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create

Receives information about reverse logistics tracking info */
func AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest, session string) (*aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse, error) {
	var resp aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
