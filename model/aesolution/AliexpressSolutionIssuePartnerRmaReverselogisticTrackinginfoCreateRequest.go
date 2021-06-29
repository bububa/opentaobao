package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create APIRequest
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create

Receives information about reverse logistics tracking info
*/
type AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest struct {
    model.Params

    // Logistic's order creation request
    logisticsOrderCreationRequest   *LogisticOrderCreationForRmaRequest 

}

func NewAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest{
    return &AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create"
}

func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) SetLogisticsOrderCreationRequest(logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest) error {
    r.logisticsOrderCreationRequest = logisticsOrderCreationRequest
    r.Set("logistics_order_creation_request", logisticsOrderCreationRequest)
    return nil
}

func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) GetLogisticsOrderCreationRequest() *LogisticOrderCreationForRmaRequest {
    return r.logisticsOrderCreationRequest
}

