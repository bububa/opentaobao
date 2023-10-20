package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushstandAPIRequest 大麦换验平台-第三方对外开放-看台接口pushStand API请求
// alibaba.damai.mev.open.pushstand
//
// pushStand
type AlibabaDamaiMevOpenPushstandAPIRequest struct {
	model.Params
	// 入参pushStandParam
	_pushStandParam *ThirdStandPushOpenParam
}

// NewAlibabaDamaiMevOpenPushstandRequest 初始化AlibabaDamaiMevOpenPushstandAPIRequest对象
func NewAlibabaDamaiMevOpenPushstandRequest() *AlibabaDamaiMevOpenPushstandAPIRequest {
	return &AlibabaDamaiMevOpenPushstandAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushstandAPIRequest) Reset() {
	r._pushStandParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushstand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushStandParam is PushStandParam Setter
// 入参pushStandParam
func (r *AlibabaDamaiMevOpenPushstandAPIRequest) SetPushStandParam(_pushStandParam *ThirdStandPushOpenParam) error {
	r._pushStandParam = _pushStandParam
	r.Set("push_stand_param", _pushStandParam)
	return nil
}

// GetPushStandParam PushStandParam Getter
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetPushStandParam() *ThirdStandPushOpenParam {
	return r._pushStandParam
}

var poolAlibabaDamaiMevOpenPushstandAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushstandRequest()
	},
}

// GetAlibabaDamaiMevOpenPushstandRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushstandAPIRequest
func GetAlibabaDamaiMevOpenPushstandAPIRequest() *AlibabaDamaiMevOpenPushstandAPIRequest {
	return poolAlibabaDamaiMevOpenPushstandAPIRequest.Get().(*AlibabaDamaiMevOpenPushstandAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushstandAPIRequest 将 AlibabaDamaiMevOpenPushstandAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushstandAPIRequest(v *AlibabaDamaiMevOpenPushstandAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushstandAPIRequest.Put(v)
}
