package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyGetAPIRequest 【商旅】查询审批单 API请求
// alitrip.btrip.corpop.apply.get
//
// 【商旅】查询审批单
type AlitripBtripCorpopApplyGetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopApplyGetRequest 初始化AlitripBtripCorpopApplyGetAPIRequest对象
func NewAlitripBtripCorpopApplyGetRequest() *AlitripBtripCorpopApplyGetAPIRequest {
	return &AlitripBtripCorpopApplyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopApplyGetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopApplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopApplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}

var poolAlitripBtripCorpopApplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopApplyGetRequest()
	},
}

// GetAlitripBtripCorpopApplyGetRequest 从 sync.Pool 获取 AlitripBtripCorpopApplyGetAPIRequest
func GetAlitripBtripCorpopApplyGetAPIRequest() *AlitripBtripCorpopApplyGetAPIRequest {
	return poolAlitripBtripCorpopApplyGetAPIRequest.Get().(*AlitripBtripCorpopApplyGetAPIRequest)
}

// ReleaseAlitripBtripCorpopApplyGetAPIRequest 将 AlitripBtripCorpopApplyGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopApplyGetAPIRequest(v *AlitripBtripCorpopApplyGetAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopApplyGetAPIRequest.Put(v)
}
