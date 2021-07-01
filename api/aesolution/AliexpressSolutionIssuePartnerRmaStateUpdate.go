package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.issue.partner.rma.state.update 
aliexpress.solution.issue.partner.rma.state.update

Receive changes in state updates for RMAs orders from after sales partners
*/
func AliexpressSolutionIssuePartnerRmaStateUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest, session string) (*aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse, error) {
    var resp aesolution.AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
