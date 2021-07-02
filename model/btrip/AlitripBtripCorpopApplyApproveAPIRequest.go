package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyApproveAPIRequest 【商旅】更新审批单状态 API请求
// alitrip.btrip.corpop.apply.approve
//
// 【商旅】更新审批单状态
type AlitripBtripCorpopApplyApproveAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiUpdateApplyRq
}

// NewAlitripBtripCorpopApplyApproveRequest 初始化AlitripBtripCorpopApplyApproveAPIRequest对象
func NewAlitripBtripCorpopApplyApproveRequest() *AlitripBtripCorpopApplyApproveAPIRequest {
	return &AlitripBtripCorpopApplyApproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyApproveAPIRequest) SetRq(_rq *OpenApiUpdateApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetRq() *OpenApiUpdateApplyRq {
	return r._rq
}
