package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalNewAPIRequest 新建审批单 API请求
// alitrip.btrip.approval.new
//
// 用户新建审批单
type AlitripBtripApprovalNewAPIRequest struct {
	model.Params
	// 申请单
	_addApplyRequest *OpenAddApplyRq
}

// NewAlitripBtripApprovalNewRequest 初始化AlitripBtripApprovalNewAPIRequest对象
func NewAlitripBtripApprovalNewRequest() *AlitripBtripApprovalNewAPIRequest {
	return &AlitripBtripApprovalNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripApprovalNewAPIRequest) Reset() {
	r._addApplyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalNewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripApprovalNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddApplyRequest is AddApplyRequest Setter
// 申请单
func (r *AlitripBtripApprovalNewAPIRequest) SetAddApplyRequest(_addApplyRequest *OpenAddApplyRq) error {
	r._addApplyRequest = _addApplyRequest
	r.Set("add_apply_request", _addApplyRequest)
	return nil
}

// GetAddApplyRequest AddApplyRequest Getter
func (r AlitripBtripApprovalNewAPIRequest) GetAddApplyRequest() *OpenAddApplyRq {
	return r._addApplyRequest
}

var poolAlitripBtripApprovalNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripApprovalNewRequest()
	},
}

// GetAlitripBtripApprovalNewRequest 从 sync.Pool 获取 AlitripBtripApprovalNewAPIRequest
func GetAlitripBtripApprovalNewAPIRequest() *AlitripBtripApprovalNewAPIRequest {
	return poolAlitripBtripApprovalNewAPIRequest.Get().(*AlitripBtripApprovalNewAPIRequest)
}

// ReleaseAlitripBtripApprovalNewAPIRequest 将 AlitripBtripApprovalNewAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripApprovalNewAPIRequest(v *AlitripBtripApprovalNewAPIRequest) {
	v.Reset()
	poolAlitripBtripApprovalNewAPIRequest.Put(v)
}
