package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopApplyApproveAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyApproveAPIRequest) SetRq(_rq *OpenApiUpdateApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopApplyApproveAPIRequest) GetRq() *OpenApiUpdateApplyRq {
	return r._rq
}

var poolAlitripBtripCorpopApplyApproveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopApplyApproveRequest()
	},
}

// GetAlitripBtripCorpopApplyApproveRequest 从 sync.Pool 获取 AlitripBtripCorpopApplyApproveAPIRequest
func GetAlitripBtripCorpopApplyApproveAPIRequest() *AlitripBtripCorpopApplyApproveAPIRequest {
	return poolAlitripBtripCorpopApplyApproveAPIRequest.Get().(*AlitripBtripCorpopApplyApproveAPIRequest)
}

// ReleaseAlitripBtripCorpopApplyApproveAPIRequest 将 AlitripBtripCorpopApplyApproveAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopApplyApproveAPIRequest(v *AlitripBtripCorpopApplyApproveAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopApplyApproveAPIRequest.Put(v)
}
