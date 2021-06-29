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
type AliexpressSolutionIssuePartnerRmaScreeningCreateRequest struct {
    model.Params
    // Screening result creation request
    _screeningResultCreationRequest   *RmaScreeningCreationRequest
}

// 初始化AliexpressSolutionIssuePartnerRmaScreeningCreateRequest对象
func NewAliexpressSolutionIssuePartnerRmaScreeningCreateRequest() *AliexpressSolutionIssuePartnerRmaScreeningCreateRequest{
    return &AliexpressSolutionIssuePartnerRmaScreeningCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.screening.create"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScreeningResultCreationRequest Setter
// Screening result creation request
func (r *AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) SetScreeningResultCreationRequest(_screeningResultCreationRequest *RmaScreeningCreationRequest) error {
    r._screeningResultCreationRequest = _screeningResultCreationRequest
    r.Set("screening_result_creation_request", _screeningResultCreationRequest)
    return nil
}

// ScreeningResultCreationRequest Getter
func (r AliexpressSolutionIssuePartnerRmaScreeningCreateRequest) GetScreeningResultCreationRequest() *RmaScreeningCreationRequest {
    return r._screeningResultCreationRequest
}
