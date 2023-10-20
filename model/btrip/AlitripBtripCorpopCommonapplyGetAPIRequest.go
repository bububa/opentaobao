package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopCommonapplyGetAPIRequest 商旅审批单通用查询接口 API请求
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
type AlitripBtripCorpopCommonapplyGetAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopCommonapplyGetRequest 初始化AlitripBtripCorpopCommonapplyGetAPIRequest对象
func NewAlitripBtripCorpopCommonapplyGetRequest() *AlitripBtripCorpopCommonapplyGetAPIRequest {
	return &AlitripBtripCorpopCommonapplyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopCommonapplyGetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.commonapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripBtripCorpopCommonapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}

var poolAlitripBtripCorpopCommonapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopCommonapplyGetRequest()
	},
}

// GetAlitripBtripCorpopCommonapplyGetRequest 从 sync.Pool 获取 AlitripBtripCorpopCommonapplyGetAPIRequest
func GetAlitripBtripCorpopCommonapplyGetAPIRequest() *AlitripBtripCorpopCommonapplyGetAPIRequest {
	return poolAlitripBtripCorpopCommonapplyGetAPIRequest.Get().(*AlitripBtripCorpopCommonapplyGetAPIRequest)
}

// ReleaseAlitripBtripCorpopCommonapplyGetAPIRequest 将 AlitripBtripCorpopCommonapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopCommonapplyGetAPIRequest(v *AlitripBtripCorpopCommonapplyGetAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopCommonapplyGetAPIRequest.Put(v)
}
