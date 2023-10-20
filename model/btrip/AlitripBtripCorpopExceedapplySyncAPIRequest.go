package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopExceedapplySyncAPIRequest 第三方超标审批结果回传 API请求
// alitrip.btrip.corpop.exceedapply.sync
//
// 第三方审批单推送到企业后，企业审批结束，将审批结果回传给阿里商旅
type AlitripBtripCorpopExceedapplySyncAPIRequest struct {
	model.Params
	// 入参
	_rq *BtripExceedApplyRq
}

// NewAlitripBtripCorpopExceedapplySyncRequest 初始化AlitripBtripCorpopExceedapplySyncAPIRequest对象
func NewAlitripBtripCorpopExceedapplySyncRequest() *AlitripBtripCorpopExceedapplySyncAPIRequest {
	return &AlitripBtripCorpopExceedapplySyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopExceedapplySyncAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.exceedapply.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopExceedapplySyncAPIRequest) SetRq(_rq *BtripExceedApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetRq() *BtripExceedApplyRq {
	return r._rq
}

var poolAlitripBtripCorpopExceedapplySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopExceedapplySyncRequest()
	},
}

// GetAlitripBtripCorpopExceedapplySyncRequest 从 sync.Pool 获取 AlitripBtripCorpopExceedapplySyncAPIRequest
func GetAlitripBtripCorpopExceedapplySyncAPIRequest() *AlitripBtripCorpopExceedapplySyncAPIRequest {
	return poolAlitripBtripCorpopExceedapplySyncAPIRequest.Get().(*AlitripBtripCorpopExceedapplySyncAPIRequest)
}

// ReleaseAlitripBtripCorpopExceedapplySyncAPIRequest 将 AlitripBtripCorpopExceedapplySyncAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopExceedapplySyncAPIRequest(v *AlitripBtripCorpopExceedapplySyncAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopExceedapplySyncAPIRequest.Put(v)
}
