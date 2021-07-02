package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AddApplyRequest Setter
// 申请单
func (r *AlitripBtripApprovalModifyAPIRequest) SetAddApplyRequest(_addApplyRequest *OpenApiNewApplyRq) error {
	r._addApplyRequest = _addApplyRequest
	r.Set("add_apply_request", _addApplyRequest)
	return nil
}

// Get AddApplyRequest Getter
func (r AlitripBtripApprovalModifyAPIRequest) GetAddApplyRequest() *OpenApiNewApplyRq {
	return r._addApplyRequest
}
