package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripapprovalnewAPIRequest 新建审批单 API请求
// alitrip.btrip.approval.new
//
// 用户新建审批单
type AlitripbtripapprovalnewAPIRequest struct {
	model.Params
	// 申请单
	_addApplyRequest *OpenAddApplyRq
}

// NewAlitripbtripapprovalnewRequest 初始化AlitripbtripapprovalnewAPIRequest对象
func NewAlitripbtripapprovalnewRequest() *AlitripbtripapprovalnewAPIRequest {
	return &AlitripbtripapprovalnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripapprovalnewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.approval.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripapprovalnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripapprovalnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddApplyRequest is AddApplyRequest Setter
// 申请单
func (r *AlitripbtripapprovalnewAPIRequest) SetAddApplyRequest(_addApplyRequest *OpenAddApplyRq) error {
	r._addApplyRequest = _addApplyRequest
	r.Set("add_apply_request", _addApplyRequest)
	return nil
}

// GetAddApplyRequest AddApplyRequest Getter
func (r AlitripbtripapprovalnewAPIRequest) GetAddApplyRequest() *OpenAddApplyRq {
	return r._addApplyRequest
}
