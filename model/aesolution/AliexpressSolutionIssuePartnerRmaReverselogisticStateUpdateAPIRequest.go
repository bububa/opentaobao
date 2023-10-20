package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest aliexpress.solution.issue.partner.rma.reverselogistic.state.update API请求
// aliexpress.solution.issue.partner.rma.reverselogistic.state.update
//
// Updates the reverse logistics state for after sales services
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest struct {
	model.Params
	// Logistic order state update request
	_logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest
}

// NewAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest 初始化AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest对象
func NewAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest {
	return &AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) Reset() {
	r._logisticOrderStateUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.reverselogistic.state.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticOrderStateUpdateRequest is LogisticOrderStateUpdateRequest Setter
// Logistic order state update request
func (r *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) SetLogisticOrderStateUpdateRequest(_logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest) error {
	r._logisticOrderStateUpdateRequest = _logisticOrderStateUpdateRequest
	r.Set("logistic_order_state_update_request", _logisticOrderStateUpdateRequest)
	return nil
}

// GetLogisticOrderStateUpdateRequest LogisticOrderStateUpdateRequest Getter
func (r AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) GetLogisticOrderStateUpdateRequest() *LogisticOrderStateUpdateRequest {
	return r._logisticOrderStateUpdateRequest
}

var poolAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest()
	},
}

// GetAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateRequest 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest
func GetAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest() *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest {
	return poolAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest.Get().(*AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest)
}

// ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest 将 AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest(v *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIRequest.Put(v)
}
