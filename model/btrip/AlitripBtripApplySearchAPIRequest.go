package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApplySearchAPIRequest 搜索审批单 API请求
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
type AlitripBtripApplySearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSearchRq
}

// NewAlitripBtripApplySearchRequest 初始化AlitripBtripApplySearchAPIRequest对象
func NewAlitripBtripApplySearchRequest() *AlitripBtripApplySearchAPIRequest {
	return &AlitripBtripApplySearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripApplySearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApplySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.apply.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApplySearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripApplySearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripApplySearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripApplySearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}

var poolAlitripBtripApplySearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripApplySearchRequest()
	},
}

// GetAlitripBtripApplySearchRequest 从 sync.Pool 获取 AlitripBtripApplySearchAPIRequest
func GetAlitripBtripApplySearchAPIRequest() *AlitripBtripApplySearchAPIRequest {
	return poolAlitripBtripApplySearchAPIRequest.Get().(*AlitripBtripApplySearchAPIRequest)
}

// ReleaseAlitripBtripApplySearchAPIRequest 将 AlitripBtripApplySearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripApplySearchAPIRequest(v *AlitripBtripApplySearchAPIRequest) {
	v.Reset()
	poolAlitripBtripApplySearchAPIRequest.Put(v)
}
