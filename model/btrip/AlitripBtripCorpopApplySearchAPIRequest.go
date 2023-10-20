package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplySearchAPIRequest 【商旅】搜索审批单列表 API请求
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
type AlitripBtripCorpopApplySearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopApplySearchRequest 初始化AlitripBtripCorpopApplySearchAPIRequest对象
func NewAlitripBtripCorpopApplySearchRequest() *AlitripBtripCorpopApplySearchAPIRequest {
	return &AlitripBtripCorpopApplySearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopApplySearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplySearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopApplySearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplySearchAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopApplySearchAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}

var poolAlitripBtripCorpopApplySearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopApplySearchRequest()
	},
}

// GetAlitripBtripCorpopApplySearchRequest 从 sync.Pool 获取 AlitripBtripCorpopApplySearchAPIRequest
func GetAlitripBtripCorpopApplySearchAPIRequest() *AlitripBtripCorpopApplySearchAPIRequest {
	return poolAlitripBtripCorpopApplySearchAPIRequest.Get().(*AlitripBtripCorpopApplySearchAPIRequest)
}

// ReleaseAlitripBtripCorpopApplySearchAPIRequest 将 AlitripBtripCorpopApplySearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopApplySearchAPIRequest(v *AlitripBtripCorpopApplySearchAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopApplySearchAPIRequest.Put(v)
}
