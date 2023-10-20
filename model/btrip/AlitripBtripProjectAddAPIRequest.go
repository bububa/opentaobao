package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectAddAPIRequest 添加项目 API请求
// alitrip.btrip.project.add
//
// 添加项目
type AlitripBtripProjectAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripBtripProjectAddRequest 初始化AlitripBtripProjectAddAPIRequest对象
func NewAlitripBtripProjectAddRequest() *AlitripBtripProjectAddAPIRequest {
	return &AlitripBtripProjectAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripProjectAddAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripProjectAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripProjectAddAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripProjectAddAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}

var poolAlitripBtripProjectAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripProjectAddRequest()
	},
}

// GetAlitripBtripProjectAddRequest 从 sync.Pool 获取 AlitripBtripProjectAddAPIRequest
func GetAlitripBtripProjectAddAPIRequest() *AlitripBtripProjectAddAPIRequest {
	return poolAlitripBtripProjectAddAPIRequest.Get().(*AlitripBtripProjectAddAPIRequest)
}

// ReleaseAlitripBtripProjectAddAPIRequest 将 AlitripBtripProjectAddAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripProjectAddAPIRequest(v *AlitripBtripProjectAddAPIRequest) {
	v.Reset()
	poolAlitripBtripProjectAddAPIRequest.Put(v)
}
