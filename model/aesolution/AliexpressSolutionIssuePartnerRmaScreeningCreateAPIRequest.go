package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.screening.create API请求
aliexpress.solution.issue.partner.rma.screening.create

Receives information about screening results from after sales partners
*/
type AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest struct {
    model.Params
    // Screening result creation request
    _screeningResultCreationRequest   *RmaScreeningCreationRequest
}

// 初始化AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest对象
func NewAliexpressSolutionIssuePartnerRmaScreeningCreateRequest() *AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest{
    return &AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.screening.create"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScreeningResultCreationRequest Setter
// Screening result creation request
func (r *AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest) SetScreeningResultCreationRequest(_screeningResultCreationRequest *RmaScreeningCreationRequest) error {
    r._screeningResultCreationRequest = _screeningResultCreationRequest
    r.Set("screening_result_creation_request", _screeningResultCreationRequest)
    return nil
}

// ScreeningResultCreationRequest Getter
func (r AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest) GetScreeningResultCreationRequest() *RmaScreeningCreationRequest {
    return r._screeningResultCreationRequest
}
