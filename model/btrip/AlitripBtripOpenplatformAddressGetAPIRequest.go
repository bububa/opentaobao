package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenplatformAddressGetAPIRequest 【商旅】开放平台对外页面跳转 API请求
// alitrip.btrip.openplatform.address.get
//
// 获取类目预定页跳转地址
type AlitripBtripOpenplatformAddressGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiJumpInfoRq
}

// NewAlitripBtripOpenplatformAddressGetRequest 初始化AlitripBtripOpenplatformAddressGetAPIRequest对象
func NewAlitripBtripOpenplatformAddressGetRequest() *AlitripBtripOpenplatformAddressGetAPIRequest {
	return &AlitripBtripOpenplatformAddressGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenplatformAddressGetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.openplatform.address.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripOpenplatformAddressGetAPIRequest) SetRq(_rq *OpenApiJumpInfoRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetRq() *OpenApiJumpInfoRq {
	return r._rq
}

var poolAlitripBtripOpenplatformAddressGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenplatformAddressGetRequest()
	},
}

// GetAlitripBtripOpenplatformAddressGetRequest 从 sync.Pool 获取 AlitripBtripOpenplatformAddressGetAPIRequest
func GetAlitripBtripOpenplatformAddressGetAPIRequest() *AlitripBtripOpenplatformAddressGetAPIRequest {
	return poolAlitripBtripOpenplatformAddressGetAPIRequest.Get().(*AlitripBtripOpenplatformAddressGetAPIRequest)
}

// ReleaseAlitripBtripOpenplatformAddressGetAPIRequest 将 AlitripBtripOpenplatformAddressGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenplatformAddressGetAPIRequest(v *AlitripBtripOpenplatformAddressGetAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenplatformAddressGetAPIRequest.Put(v)
}
