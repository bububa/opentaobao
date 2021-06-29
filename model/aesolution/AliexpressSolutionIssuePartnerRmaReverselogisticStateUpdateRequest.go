package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.reverselogistic.state.update API请求
aliexpress.solution.issue.partner.rma.reverselogistic.state.update

Updates the reverse logistics state for after sales services
*/
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest struct {
    model.Params
    // Logistic order state update request
    logisticOrderStateUpdateRequest   *LogisticOrderStateUpdateRequest
}

// 初始化AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest对象
func NewAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest{
    return &AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.reverselogistic.state.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisticOrderStateUpdateRequest Setter
// Logistic order state update request
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) SetLogisticOrderStateUpdateRequest(logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest) error {
    r.logisticOrderStateUpdateRequest = logisticOrderStateUpdateRequest
    r.Set("logistic_order_state_update_request", logisticOrderStateUpdateRequest)
    return nil
}

// LogisticOrderStateUpdateRequest Getter
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest) GetLogisticOrderStateUpdateRequest() *LogisticOrderStateUpdateRequest {
    return r.logisticOrderStateUpdateRequest
}
