package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfaceAPIRequest 大麦换验平台-第三方对外开放-票面接口pushFace API请求
// alibaba.damai.mev.open.pushface
//
// pushFace
type AlibabaDamaiMevOpenPushfaceAPIRequest struct {
	model.Params
	// 入参pushFaceParam
	_pushFaceParam *ThirdTicketFacePushOpenParam
}

// NewAlibabaDamaiMevOpenPushfaceRequest 初始化AlibabaDamaiMevOpenPushfaceAPIRequest对象
func NewAlibabaDamaiMevOpenPushfaceRequest() *AlibabaDamaiMevOpenPushfaceAPIRequest {
	return &AlibabaDamaiMevOpenPushfaceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushfaceAPIRequest) Reset() {
	r._pushFaceParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushFaceParam is PushFaceParam Setter
// 入参pushFaceParam
func (r *AlibabaDamaiMevOpenPushfaceAPIRequest) SetPushFaceParam(_pushFaceParam *ThirdTicketFacePushOpenParam) error {
	r._pushFaceParam = _pushFaceParam
	r.Set("push_face_param", _pushFaceParam)
	return nil
}

// GetPushFaceParam PushFaceParam Getter
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetPushFaceParam() *ThirdTicketFacePushOpenParam {
	return r._pushFaceParam
}

var poolAlibabaDamaiMevOpenPushfaceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushfaceRequest()
	},
}

// GetAlibabaDamaiMevOpenPushfaceRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushfaceAPIRequest
func GetAlibabaDamaiMevOpenPushfaceAPIRequest() *AlibabaDamaiMevOpenPushfaceAPIRequest {
	return poolAlibabaDamaiMevOpenPushfaceAPIRequest.Get().(*AlibabaDamaiMevOpenPushfaceAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushfaceAPIRequest 将 AlibabaDamaiMevOpenPushfaceAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushfaceAPIRequest(v *AlibabaDamaiMevOpenPushfaceAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushfaceAPIRequest.Put(v)
}
