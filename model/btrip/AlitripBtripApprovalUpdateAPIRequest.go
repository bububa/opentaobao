package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalUpdateAPIRequest 更新审批单 API请求
// alitrip.btrip.approval.update
//
// 更新审批单
type AlitripBtripApprovalUpdateAPIRequest struct {
	model.Params
	// 审批请求对象
	_approveApplyRequest *OpenApproveApplyRq
}

// NewAlitripBtripApprovalUpdateRequest 初始化AlitripBtripApprovalUpdateAPIRequest对象
func NewAlitripBtripApprovalUpdateRequest() *AlitripBtripApprovalUpdateAPIRequest {
	return &AlitripBtripApprovalUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripApprovalUpdateAPIRequest) Reset() {
	r._approveApplyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripApprovalUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApproveApplyRequest is ApproveApplyRequest Setter
// 审批请求对象
func (r *AlitripBtripApprovalUpdateAPIRequest) SetApproveApplyRequest(_approveApplyRequest *OpenApproveApplyRq) error {
	r._approveApplyRequest = _approveApplyRequest
	r.Set("approve_apply_request", _approveApplyRequest)
	return nil
}

// GetApproveApplyRequest ApproveApplyRequest Getter
func (r AlitripBtripApprovalUpdateAPIRequest) GetApproveApplyRequest() *OpenApproveApplyRq {
	return r._approveApplyRequest
}

var poolAlitripBtripApprovalUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripApprovalUpdateRequest()
	},
}

// GetAlitripBtripApprovalUpdateRequest 从 sync.Pool 获取 AlitripBtripApprovalUpdateAPIRequest
func GetAlitripBtripApprovalUpdateAPIRequest() *AlitripBtripApprovalUpdateAPIRequest {
	return poolAlitripBtripApprovalUpdateAPIRequest.Get().(*AlitripBtripApprovalUpdateAPIRequest)
}

// ReleaseAlitripBtripApprovalUpdateAPIRequest 将 AlitripBtripApprovalUpdateAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripApprovalUpdateAPIRequest(v *AlitripBtripApprovalUpdateAPIRequest) {
	v.Reset()
	poolAlitripBtripApprovalUpdateAPIRequest.Put(v)
}
