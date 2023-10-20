package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletefaceAPIRequest 大麦换验平台-第三方对外开放-票面接口deleteFace API请求
// alibaba.damai.mev.open.deleteface
//
// deleteFace
type AlibabaDamaiMevOpenDeletefaceAPIRequest struct {
	model.Params
	// 入参deleteFaceParam
	_deleteFaceParam *TicketFaceIdOpenParam
}

// NewAlibabaDamaiMevOpenDeletefaceRequest 初始化AlibabaDamaiMevOpenDeletefaceAPIRequest对象
func NewAlibabaDamaiMevOpenDeletefaceRequest() *AlibabaDamaiMevOpenDeletefaceAPIRequest {
	return &AlibabaDamaiMevOpenDeletefaceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenDeletefaceAPIRequest) Reset() {
	r._deleteFaceParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletefaceAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletefaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeletefaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteFaceParam is DeleteFaceParam Setter
// 入参deleteFaceParam
func (r *AlibabaDamaiMevOpenDeletefaceAPIRequest) SetDeleteFaceParam(_deleteFaceParam *TicketFaceIdOpenParam) error {
	r._deleteFaceParam = _deleteFaceParam
	r.Set("delete_face_param", _deleteFaceParam)
	return nil
}

// GetDeleteFaceParam DeleteFaceParam Getter
func (r AlibabaDamaiMevOpenDeletefaceAPIRequest) GetDeleteFaceParam() *TicketFaceIdOpenParam {
	return r._deleteFaceParam
}

var poolAlibabaDamaiMevOpenDeletefaceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenDeletefaceRequest()
	},
}

// GetAlibabaDamaiMevOpenDeletefaceRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletefaceAPIRequest
func GetAlibabaDamaiMevOpenDeletefaceAPIRequest() *AlibabaDamaiMevOpenDeletefaceAPIRequest {
	return poolAlibabaDamaiMevOpenDeletefaceAPIRequest.Get().(*AlibabaDamaiMevOpenDeletefaceAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenDeletefaceAPIRequest 将 AlibabaDamaiMevOpenDeletefaceAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletefaceAPIRequest(v *AlibabaDamaiMevOpenDeletefaceAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletefaceAPIRequest.Put(v)
}
