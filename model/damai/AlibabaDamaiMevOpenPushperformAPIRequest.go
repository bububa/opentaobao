package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushperformAPIRequest 大麦换验平台-第三方对外开放-场次接口pushPerform API请求
// alibaba.damai.mev.open.pushperform
//
// pushPerform
type AlibabaDamaiMevOpenPushperformAPIRequest struct {
	model.Params
	// 入参pushPerformParam
	_pushPerformParam *ThirdPerformPushOpenParam
}

// NewAlibabaDamaiMevOpenPushperformRequest 初始化AlibabaDamaiMevOpenPushperformAPIRequest对象
func NewAlibabaDamaiMevOpenPushperformRequest() *AlibabaDamaiMevOpenPushperformAPIRequest {
	return &AlibabaDamaiMevOpenPushperformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushperformAPIRequest) Reset() {
	r._pushPerformParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushperform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushPerformParam is PushPerformParam Setter
// 入参pushPerformParam
func (r *AlibabaDamaiMevOpenPushperformAPIRequest) SetPushPerformParam(_pushPerformParam *ThirdPerformPushOpenParam) error {
	r._pushPerformParam = _pushPerformParam
	r.Set("push_perform_param", _pushPerformParam)
	return nil
}

// GetPushPerformParam PushPerformParam Getter
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetPushPerformParam() *ThirdPerformPushOpenParam {
	return r._pushPerformParam
}

var poolAlibabaDamaiMevOpenPushperformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushperformRequest()
	},
}

// GetAlibabaDamaiMevOpenPushperformRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushperformAPIRequest
func GetAlibabaDamaiMevOpenPushperformAPIRequest() *AlibabaDamaiMevOpenPushperformAPIRequest {
	return poolAlibabaDamaiMevOpenPushperformAPIRequest.Get().(*AlibabaDamaiMevOpenPushperformAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushperformAPIRequest 将 AlibabaDamaiMevOpenPushperformAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushperformAPIRequest(v *AlibabaDamaiMevOpenPushperformAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushperformAPIRequest.Put(v)
}
