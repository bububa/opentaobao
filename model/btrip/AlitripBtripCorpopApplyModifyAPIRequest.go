package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyModifyAPIRequest 【商旅】修改出差审批单（行程） API请求
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
type AlitripBtripCorpopApplyModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiApplyRq
}

// NewAlitripBtripCorpopApplyModifyRequest 初始化AlitripBtripCorpopApplyModifyAPIRequest对象
func NewAlitripBtripCorpopApplyModifyRequest() *AlitripBtripCorpopApplyModifyAPIRequest {
	return &AlitripBtripCorpopApplyModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopApplyModifyAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyModifyAPIRequest) SetRq(_rq *OpenApiApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetRq() *OpenApiApplyRq {
	return r._rq
}

var poolAlitripBtripCorpopApplyModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopApplyModifyRequest()
	},
}

// GetAlitripBtripCorpopApplyModifyRequest 从 sync.Pool 获取 AlitripBtripCorpopApplyModifyAPIRequest
func GetAlitripBtripCorpopApplyModifyAPIRequest() *AlitripBtripCorpopApplyModifyAPIRequest {
	return poolAlitripBtripCorpopApplyModifyAPIRequest.Get().(*AlitripBtripCorpopApplyModifyAPIRequest)
}

// ReleaseAlitripBtripCorpopApplyModifyAPIRequest 将 AlitripBtripCorpopApplyModifyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopApplyModifyAPIRequest(v *AlitripBtripCorpopApplyModifyAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopApplyModifyAPIRequest.Put(v)
}
