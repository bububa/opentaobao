package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create API请求
// aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
//
// Receives information about reverse logistics tracking info
type AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest struct {
	model.Params
	// Logistic's order creation request
	_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest
}

// NewAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest 初始化AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest对象
func NewAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest {
	return &AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LogisticsOrderCreationRequest Setter
// Logistic's order creation request
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) SetLogisticsOrderCreationRequest(_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest) error {
	r._logisticsOrderCreationRequest = _logisticsOrderCreationRequest
	r.Set("logistics_order_creation_request", _logisticsOrderCreationRequest)
	return nil
}

// Get LogisticsOrderCreationRequest Getter
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetLogisticsOrderCreationRequest() *LogisticOrderCreationForRmaRequest {
	return r._logisticsOrderCreationRequest
}
