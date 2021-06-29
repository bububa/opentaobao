package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.screening.create APIRequest
aliexpress.solution.issue.partner.rma.screening.create

Receives information about screening results from after sales partners
*/
type AliexpressSolutionIssuePartnerRmaScreeningCreateRequest struct {
    model.Params

    // Screening result creation request
    screeningResultCreationRequest   *RmaScreeningCreationRequest 

}

func NewAliexpressSolutionIssuePartnerRmaScreeningCreateRequest() *AliexpressSolutionIssuePartnerRmaScreeningCreateRequest{
    return &AliexpressSolutionIssuePartnerRmaScreeningCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.screening.create"
}

func (r AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) SetScreeningResultCreationRequest(screeningResultCreationRequest *RmaScreeningCreationRequest) error {
    r.screeningResultCreationRequest = screeningResultCreationRequest
    r.Set("screening_result_creation_request", screeningResultCreationRequest)
    return nil
}

func (r AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) GetScreeningResultCreationRequest() *RmaScreeningCreationRequest {
    return r.screeningResultCreationRequest
}

