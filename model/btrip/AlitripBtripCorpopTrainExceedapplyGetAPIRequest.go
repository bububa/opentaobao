package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopTrainExceedapplyGetAPIRequest 商旅火车票第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
type AlitripBtripCorpopTrainExceedapplyGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopTrainExceedapplyGetRequest 初始化AlitripBtripCorpopTrainExceedapplyGetAPIRequest对象
func NewAlitripBtripCorpopTrainExceedapplyGetRequest() *AlitripBtripCorpopTrainExceedapplyGetAPIRequest {
	return &AlitripBtripCorpopTrainExceedapplyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopTrainExceedapplyGetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.train.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopTrainExceedapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}

var poolAlitripBtripCorpopTrainExceedapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopTrainExceedapplyGetRequest()
	},
}

// GetAlitripBtripCorpopTrainExceedapplyGetRequest 从 sync.Pool 获取 AlitripBtripCorpopTrainExceedapplyGetAPIRequest
func GetAlitripBtripCorpopTrainExceedapplyGetAPIRequest() *AlitripBtripCorpopTrainExceedapplyGetAPIRequest {
	return poolAlitripBtripCorpopTrainExceedapplyGetAPIRequest.Get().(*AlitripBtripCorpopTrainExceedapplyGetAPIRequest)
}

// ReleaseAlitripBtripCorpopTrainExceedapplyGetAPIRequest 将 AlitripBtripCorpopTrainExceedapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopTrainExceedapplyGetAPIRequest(v *AlitripBtripCorpopTrainExceedapplyGetAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopTrainExceedapplyGetAPIRequest.Put(v)
}
