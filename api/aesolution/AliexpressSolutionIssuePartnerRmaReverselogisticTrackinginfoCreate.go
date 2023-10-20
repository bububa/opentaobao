package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionissuepartnerrmareverselogistictrackinginfocreate aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
// aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
//
// Receives information about reverse logistics tracking info
func Aliexpresssolutionissuepartnerrmareverselogistictrackinginfocreate(clt *core.SDKClient, req *aesolution.AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest, session string) (*aesolution.AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
