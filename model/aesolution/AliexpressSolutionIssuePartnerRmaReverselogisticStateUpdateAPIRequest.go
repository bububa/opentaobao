package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest aliexpress.solution.issue.partner.rma.reverselogistic.state.update API请求
// aliexpress.solution.issue.partner.rma.reverselogistic.state.update
//
// Updates the reverse logistics state for after sales services
type AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest struct {
	model.Params
	// Logistic order state update request
	_logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest
}

// NewAliexpresssolutionissuepartnerrmareverselogisticstateupdateRequest 初始化AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest对象
func NewAliexpresssolutionissuepartnerrmareverselogisticstateupdateRequest() *AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest {
	return &AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.reverselogistic.state.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticOrderStateUpdateRequest is LogisticOrderStateUpdateRequest Setter
// Logistic order state update request
func (r *AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest) SetLogisticOrderStateUpdateRequest(_logisticOrderStateUpdateRequest *LogisticOrderStateUpdateRequest) error {
	r._logisticOrderStateUpdateRequest = _logisticOrderStateUpdateRequest
	r.Set("logistic_order_state_update_request", _logisticOrderStateUpdateRequest)
	return nil
}

// GetLogisticOrderStateUpdateRequest LogisticOrderStateUpdateRequest Getter
func (r AliexpresssolutionissuepartnerrmareverselogisticstateupdateAPIRequest) GetLogisticOrderStateUpdateRequest() *LogisticOrderStateUpdateRequest {
	return r._logisticOrderStateUpdateRequest
}
