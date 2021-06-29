package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.reverselogistic.state.update APIRequest
aliexpress.solution.issue.partner.rma.reverselogistic.state.update

Updates the reverse logistics state for after sales services
*/
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest struct {
    model.Params

    // Logistic order state update request
    logisticOrderStateUpdateRequest   *LogisticOrderStateUpdateRequest 

}

func NewAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest{
    return &AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.reverselogistic.state.update"
}

func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) SetLogisticOrderStateUpdateRequest(logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest) error {
    r.logisticOrderStateUpdateRequest = logisticOrderStateUpdateRequest
    r.Set("logistic_order_state_update_request", logisticOrderStateUpdateRequest)
    return nil
}

func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) GetLogisticOrderStateUpdateRequest() *LogisticOrderStateUpdateRequest {
    return r.logisticOrderStateUpdateRequest
}

