package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.state.update API请求
aliexpress.solution.issue.partner.rma.state.update

Receive changes in state updates for RMAs orders from after sales partners
*/
type AliexpressSolutionIssuePartnerRmaStateUpdateRequest struct {
    model.Params
    // RMA's order state update request
    _rmaStateUpdateRequest   *RmaStateUpdateRequest
}

// 初始化AliexpressSolutionIssuePartnerRmaStateUpdateRequest对象
func NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaStateUpdateRequest{
    return &AliexpressSolutionIssuePartnerRmaStateUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaStateUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.state.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaStateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RmaStateUpdateRequest Setter
// RMA's order state update request
func (r *AliexpressSolutionIssuePartnerRmaStateUpdateRequest) SetRmaStateUpdateRequest(_rmaStateUpdateRequest *RmaStateUpdateRequest) error {
    r._rmaStateUpdateRequest = _rmaStateUpdateRequest
    r.Set("rma_state_update_request", _rmaStateUpdateRequest)
    return nil
}

// RmaStateUpdateRequest Getter
func (r AliexpressSolutionIssuePartnerRmaStateUpdateRequest) GetRmaStateUpdateRequest() *RmaStateUpdateRequest {
    return r._rmaStateUpdateRequest
}
