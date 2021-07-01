package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripApprovalUpdateAPIRequest
更新审批单 API请求
alitrip.btrip.approval.update

更新审批单 */
type AlitripBtripApprovalUpdateAPIRequest struct {
	model.Params
	// 审批请求对象
	_approveApplyRequest *OpenApproveApplyRq
}

// NewAlitripBtripApprovalUpdateRequest 初始化AlitripBtripApprovalUpdateAPIRequest对象
func NewAlitripBtripApprovalUpdateRequest() *AlitripBtripApprovalUpdateAPIRequest {
	return &AlitripBtripApprovalUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ApproveApplyRequest Setter
// 审批请求对象
func (r *AlitripBtripApprovalUpdateAPIRequest) SetApproveApplyRequest(_approveApplyRequest *OpenApproveApplyRq) error {
	r._approveApplyRequest = _approveApplyRequest
	r.Set("approve_apply_request", _approveApplyRequest)
	return nil
}

// Get ApproveApplyRequest Getter
func (r AlitripBtripApprovalUpdateAPIRequest) GetApproveApplyRequest() *OpenApproveApplyRq {
	return r._approveApplyRequest
}
