package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyAddAPIRequest 【商旅】isv添加审批单 API请求
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
type AlitripBtripCorpopApplyAddAPIRequest struct {
	model.Params
	// 请求参数
	_rq *OpenApiApplyRq
}

// NewAlitripBtripCorpopApplyAddRequest 初始化AlitripBtripCorpopApplyAddAPIRequest对象
func NewAlitripBtripCorpopApplyAddRequest() *AlitripBtripCorpopApplyAddAPIRequest {
	return &AlitripBtripCorpopApplyAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopApplyAddAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopApplyAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求参数
func (r *AlitripBtripCorpopApplyAddAPIRequest) SetRq(_rq *OpenApiApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopApplyAddAPIRequest) GetRq() *OpenApiApplyRq {
	return r._rq
}

var poolAlitripBtripCorpopApplyAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopApplyAddRequest()
	},
}

// GetAlitripBtripCorpopApplyAddRequest 从 sync.Pool 获取 AlitripBtripCorpopApplyAddAPIRequest
func GetAlitripBtripCorpopApplyAddAPIRequest() *AlitripBtripCorpopApplyAddAPIRequest {
	return poolAlitripBtripCorpopApplyAddAPIRequest.Get().(*AlitripBtripCorpopApplyAddAPIRequest)
}

// ReleaseAlitripBtripCorpopApplyAddAPIRequest 将 AlitripBtripCorpopApplyAddAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopApplyAddAPIRequest(v *AlitripBtripCorpopApplyAddAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopApplyAddAPIRequest.Put(v)
}
