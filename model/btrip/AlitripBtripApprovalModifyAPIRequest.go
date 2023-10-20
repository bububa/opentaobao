package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripapprovalmodifyAPIRequest 修改审批单 API请求
// alitrip.btrip.approval.modify
//
// 修改审批单
type AlitripbtripapprovalmodifyAPIRequest struct {
	model.Params
	// 申请单
	_addApplyRequest *OpenApiNewApplyRq
}

// NewAlitripbtripapprovalmodifyRequest 初始化AlitripbtripapprovalmodifyAPIRequest对象
func NewAlitripbtripapprovalmodifyRequest() *AlitripbtripapprovalmodifyAPIRequest {
	return &AlitripbtripapprovalmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripapprovalmodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripapprovalmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripapprovalmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddApplyRequest is AddApplyRequest Setter
// 申请单
func (r *AlitripbtripapprovalmodifyAPIRequest) SetAddApplyRequest(_addApplyRequest *OpenApiNewApplyRq) error {
	r._addApplyRequest = _addApplyRequest
	r.Set("add_apply_request", _addApplyRequest)
	return nil
}

// GetAddApplyRequest AddApplyRequest Getter
func (r AlitripbtripapprovalmodifyAPIRequest) GetAddApplyRequest() *OpenApiNewApplyRq {
	return r._addApplyRequest
}
