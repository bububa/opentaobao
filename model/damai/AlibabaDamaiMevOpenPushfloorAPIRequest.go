package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfloorAPIRequest 大麦换验平台-第三方对外开放-楼层接口pushFloor API请求
// alibaba.damai.mev.open.pushfloor
//
// pushFloor
type AlibabaDamaiMevOpenPushfloorAPIRequest struct {
	model.Params
	// 入参pushFloorParam
	_pushFloorParam *ThirdFloorPushOpenParam
}

// NewAlibabaDamaiMevOpenPushfloorRequest 初始化AlibabaDamaiMevOpenPushfloorAPIRequest对象
func NewAlibabaDamaiMevOpenPushfloorRequest() *AlibabaDamaiMevOpenPushfloorAPIRequest {
	return &AlibabaDamaiMevOpenPushfloorAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushfloorAPIRequest) Reset() {
	r._pushFloorParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushfloor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushFloorParam is PushFloorParam Setter
// 入参pushFloorParam
func (r *AlibabaDamaiMevOpenPushfloorAPIRequest) SetPushFloorParam(_pushFloorParam *ThirdFloorPushOpenParam) error {
	r._pushFloorParam = _pushFloorParam
	r.Set("push_floor_param", _pushFloorParam)
	return nil
}

// GetPushFloorParam PushFloorParam Getter
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetPushFloorParam() *ThirdFloorPushOpenParam {
	return r._pushFloorParam
}

var poolAlibabaDamaiMevOpenPushfloorAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushfloorRequest()
	},
}

// GetAlibabaDamaiMevOpenPushfloorRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushfloorAPIRequest
func GetAlibabaDamaiMevOpenPushfloorAPIRequest() *AlibabaDamaiMevOpenPushfloorAPIRequest {
	return poolAlibabaDamaiMevOpenPushfloorAPIRequest.Get().(*AlibabaDamaiMevOpenPushfloorAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushfloorAPIRequest 将 AlibabaDamaiMevOpenPushfloorAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushfloorAPIRequest(v *AlibabaDamaiMevOpenPushfloorAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushfloorAPIRequest.Put(v)
}
