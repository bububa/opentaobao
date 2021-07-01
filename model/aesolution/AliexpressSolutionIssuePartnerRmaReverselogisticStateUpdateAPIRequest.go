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
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest struct {
    model.Params
    // Logistic order state update request
    _logisticOrderStateUpdateRequest   *LogisticOrderStateUpdateRequest
}

// 初始化AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest对象
func NewAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest{
    return &AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.reverselogistic.state.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisticOrderStateUpdateRequest Setter
// Logistic order state update request
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) SetLogisticOrderStateUpdateRequest(_logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest) error {
    r._logisticOrderStateUpdateRequest = _logisticOrderStateUpdateRequest
    r.Set("logistic_order_state_update_request", _logisticOrderStateUpdateRequest)
    return nil
}

// LogisticOrderStateUpdateRequest Getter
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetLogisticOrderStateUpdateRequest() *LogisticOrderStateUpdateRequest {
    return r._logisticOrderStateUpdateRequest
}
