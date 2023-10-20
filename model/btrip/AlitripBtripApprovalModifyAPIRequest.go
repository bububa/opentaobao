package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalModifyAPIRequest 修改审批单 API请求
// alitrip.btrip.approval.modify
//
// 修改审批单
type AlitripBtripApprovalModifyAPIRequest struct {
	model.Params
	// 申请单
	_addApplyRequest *OpenApiNewApplyRq
}

// NewAlitripBtripApprovalModifyRequest 初始化AlitripBtripApprovalModifyAPIRequest对象
func NewAlitripBtripApprovalModifyRequest() *AlitripBtripApprovalModifyAPIRequest {
	return &AlitripBtripApprovalModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripApprovalModifyAPIRequest) Reset() {
	r._addApplyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripApprovalModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddApplyRequest is AddApplyRequest Setter
// 申请单
func (r *AlitripBtripApprovalModifyAPIRequest) SetAddApplyRequest(_addApplyRequest *OpenApiNewApplyRq) error {
	r._addApplyRequest = _addApplyRequest
	r.Set("add_apply_request", _addApplyRequest)
	return nil
}

// GetAddApplyRequest AddApplyRequest Getter
func (r AlitripBtripApprovalModifyAPIRequest) GetAddApplyRequest() *OpenApiNewApplyRq {
	return r._addApplyRequest
}

var poolAlitripBtripApprovalModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripApprovalModifyRequest()
	},
}

// GetAlitripBtripApprovalModifyRequest 从 sync.Pool 获取 AlitripBtripApprovalModifyAPIRequest
func GetAlitripBtripApprovalModifyAPIRequest() *AlitripBtripApprovalModifyAPIRequest {
	return poolAlitripBtripApprovalModifyAPIRequest.Get().(*AlitripBtripApprovalModifyAPIRequest)
}

// ReleaseAlitripBtripApprovalModifyAPIRequest 将 AlitripBtripApprovalModifyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripApprovalModifyAPIRequest(v *AlitripBtripApprovalModifyAPIRequest) {
	v.Reset()
	poolAlitripBtripApprovalModifyAPIRequest.Put(v)
}
