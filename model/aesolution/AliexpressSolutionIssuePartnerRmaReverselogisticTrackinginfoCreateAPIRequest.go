package aesolution

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) Reset() {
	r._logisticsOrderCreationRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsOrderCreationRequest is LogisticsOrderCreationRequest Setter
// Logistic&#39;s order creation request
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) SetLogisticsOrderCreationRequest(_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest) error {
	r._logisticsOrderCreationRequest = _logisticsOrderCreationRequest
	r.Set("logistics_order_creation_request", _logisticsOrderCreationRequest)
	return nil
}

// GetLogisticsOrderCreationRequest LogisticsOrderCreationRequest Getter
func (r AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) GetLogisticsOrderCreationRequest() *LogisticOrderCreationForRmaRequest {
	return r._logisticsOrderCreationRequest
}

var poolAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest()
	},
}

// GetAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateRequest 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest
func GetAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest {
	return poolAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest.Get().(*AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest)
}

// ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest 将 AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest(v *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest.Put(v)
}
