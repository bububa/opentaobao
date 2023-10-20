package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionissuepartnerrmastateupdateAPIRequest aliexpress.solution.issue.partner.rma.state.update API请求
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
type AliexpresssolutionissuepartnerrmastateupdateAPIRequest struct {
	model.Params
	// RMA's order state update request
	_rmaStateUpdateRequest *RmaStateUpdateRequest
}

// NewAliexpresssolutionissuepartnerrmastateupdateRequest 初始化AliexpresssolutionissuepartnerrmastateupdateAPIRequest对象
func NewAliexpresssolutionissuepartnerrmastateupdateRequest() *AliexpresssolutionissuepartnerrmastateupdateAPIRequest {
	return &AliexpresssolutionissuepartnerrmastateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionissuepartnerrmastateupdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.state.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionissuepartnerrmastateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionissuepartnerrmastateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRmaStateUpdateRequest is RmaStateUpdateRequest Setter
// RMA&#39;s order state update request
func (r *AliexpresssolutionissuepartnerrmastateupdateAPIRequest) SetRmaStateUpdateRequest(_rmaStateUpdateRequest *RmaStateUpdateRequest) error {
	r._rmaStateUpdateRequest = _rmaStateUpdateRequest
	r.Set("rma_state_update_request", _rmaStateUpdateRequest)
	return nil
}

// GetRmaStateUpdateRequest RmaStateUpdateRequest Getter
func (r AliexpresssolutionissuepartnerrmastateupdateAPIRequest) GetRmaStateUpdateRequest() *RmaStateUpdateRequest {
	return r._rmaStateUpdateRequest
}
