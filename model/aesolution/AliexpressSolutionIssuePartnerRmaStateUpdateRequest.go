package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.state.update APIRequest
aliexpress.solution.issue.partner.rma.state.update

Receive changes in state updates for RMAs orders from after sales partners
*/
type AliexpressSolutionIssuePartnerRmaStateUpdateRequest struct {
    model.Params

    // RMA's order state update request
    rmaStateUpdateRequest   *RmaStateUpdateRequest 

}

func NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaStateUpdateRequest{
    return &AliexpressSolutionIssuePartnerRmaStateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionIssuePartnerRmaStateUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.state.update"
}

func (r AliexpressSolutionIssuePartnerRmaStateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionIssuePartnerRmaStateUpdateRequest) SetRmaStateUpdateRequest(rmaStateUpdateRequest *RmaStateUpdateRequest) error {
    r.rmaStateUpdateRequest = rmaStateUpdateRequest
    r.Set("rma_state_update_request", rmaStateUpdateRequest)
    return nil
}

func (r AliexpressSolutionIssuePartnerRmaStateUpdateRequest) GetRmaStateUpdateRequest() *RmaStateUpdateRequest {
    return r.rmaStateUpdateRequest
}

