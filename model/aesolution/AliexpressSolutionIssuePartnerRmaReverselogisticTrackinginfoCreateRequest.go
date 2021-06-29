package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create API请求
aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create

Receives information about reverse logistics tracking info
*/
type AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest struct {
    model.Params
    // Logistic's order creation request
    _logisticsOrderCreationRequest   *LogisticOrderCreationForRmaRequest
}

// 初始化AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest对象
func NewAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest{
    return &AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) GetApiMethodName() string {
    return "aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisticsOrderCreationRequest Setter
// Logistic's order creation request
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) SetLogisticsOrderCreationRequest(_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest) error {
    r._logisticsOrderCreationRequest = _logisticsOrderCreationRequest
    r.Set("logistics_order_creation_request", _logisticsOrderCreationRequest)
    return nil
}

// LogisticsOrderCreationRequest Getter
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest) GetLogisticsOrderCreationRequest() *LogisticOrderCreationForRmaRequest {
    return r._logisticsOrderCreationRequest
}
