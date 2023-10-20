package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectModifyAPIRequest 变更项目 API请求
// alitrip.btrip.project.modify
//
// 变更项目
type AlitripBtripProjectModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripBtripProjectModifyRequest 初始化AlitripBtripProjectModifyAPIRequest对象
func NewAlitripBtripProjectModifyRequest() *AlitripBtripProjectModifyAPIRequest {
	return &AlitripBtripProjectModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripProjectModifyAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripProjectModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripProjectModifyAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripProjectModifyAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}

var poolAlitripBtripProjectModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripProjectModifyRequest()
	},
}

// GetAlitripBtripProjectModifyRequest 从 sync.Pool 获取 AlitripBtripProjectModifyAPIRequest
func GetAlitripBtripProjectModifyAPIRequest() *AlitripBtripProjectModifyAPIRequest {
	return poolAlitripBtripProjectModifyAPIRequest.Get().(*AlitripBtripProjectModifyAPIRequest)
}

// ReleaseAlitripBtripProjectModifyAPIRequest 将 AlitripBtripProjectModifyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripProjectModifyAPIRequest(v *AlitripBtripProjectModifyAPIRequest) {
	v.Reset()
	poolAlitripBtripProjectModifyAPIRequest.Put(v)
}
