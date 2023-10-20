package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopFlightExceedapplyGetAPIRequest 商旅机票第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
type AlitripBtripCorpopFlightExceedapplyGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopFlightExceedapplyGetRequest 初始化AlitripBtripCorpopFlightExceedapplyGetAPIRequest对象
func NewAlitripBtripCorpopFlightExceedapplyGetRequest() *AlitripBtripCorpopFlightExceedapplyGetAPIRequest {
	return &AlitripBtripCorpopFlightExceedapplyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopFlightExceedapplyGetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.flight.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopFlightExceedapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}

var poolAlitripBtripCorpopFlightExceedapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopFlightExceedapplyGetRequest()
	},
}

// GetAlitripBtripCorpopFlightExceedapplyGetRequest 从 sync.Pool 获取 AlitripBtripCorpopFlightExceedapplyGetAPIRequest
func GetAlitripBtripCorpopFlightExceedapplyGetAPIRequest() *AlitripBtripCorpopFlightExceedapplyGetAPIRequest {
	return poolAlitripBtripCorpopFlightExceedapplyGetAPIRequest.Get().(*AlitripBtripCorpopFlightExceedapplyGetAPIRequest)
}

// ReleaseAlitripBtripCorpopFlightExceedapplyGetAPIRequest 将 AlitripBtripCorpopFlightExceedapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopFlightExceedapplyGetAPIRequest(v *AlitripBtripCorpopFlightExceedapplyGetAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopFlightExceedapplyGetAPIRequest.Put(v)
}
