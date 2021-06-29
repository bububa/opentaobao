package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.issue.partner.rma.reverselogistic.state.update 
aliexpress.solution.issue.partner.rma.reverselogistic.state.update

Updates the reverse logistics state for after sales services
*/
func AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest, session string) (*aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse, error) {
    var resp aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
