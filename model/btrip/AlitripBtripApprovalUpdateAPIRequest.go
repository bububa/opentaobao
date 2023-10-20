package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripapprovalupdateAPIRequest 更新审批单 API请求
// alitrip.btrip.approval.update
//
// 更新审批单
type AlitripbtripapprovalupdateAPIRequest struct {
	model.Params
	// 审批请求对象
	_approveApplyRequest *OpenApproveApplyRq
}

// NewAlitripbtripapprovalupdateRequest 初始化AlitripbtripapprovalupdateAPIRequest对象
func NewAlitripbtripapprovalupdateRequest() *AlitripbtripapprovalupdateAPIRequest {
	return &AlitripbtripapprovalupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripapprovalupdateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripapprovalupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripapprovalupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApproveApplyRequest is ApproveApplyRequest Setter
// 审批请求对象
func (r *AlitripbtripapprovalupdateAPIRequest) SetApproveApplyRequest(_approveApplyRequest *OpenApproveApplyRq) error {
	r._approveApplyRequest = _approveApplyRequest
	r.Set("approve_apply_request", _approveApplyRequest)
	return nil
}

// GetApproveApplyRequest ApproveApplyRequest Getter
func (r AlitripbtripapprovalupdateAPIRequest) GetApproveApplyRequest() *OpenApproveApplyRq {
	return r._approveApplyRequest
}
