package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteFaceelementAPIRequest 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement API请求
// alibaba.damai.mev.open.delete.faceelement
//
// deleteFaceElement
type AlibabaDamaiMevOpenDeleteFaceelementAPIRequest struct {
	model.Params
	// 入参deleteFaceElementParam
	_deleteFaceElementParam *FaceElementIdOpenParam
}

// NewAlibabaDamaiMevOpenDeleteFaceelementRequest 初始化AlibabaDamaiMevOpenDeleteFaceelementAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteFaceelementRequest() *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest {
	return &AlibabaDamaiMevOpenDeleteFaceelementAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) Reset() {
	r._deleteFaceElementParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.delete.faceelement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteFaceElementParam is DeleteFaceElementParam Setter
// 入参deleteFaceElementParam
func (r *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) SetDeleteFaceElementParam(_deleteFaceElementParam *FaceElementIdOpenParam) error {
	r._deleteFaceElementParam = _deleteFaceElementParam
	r.Set("delete_face_element_param", _deleteFaceElementParam)
	return nil
}

// GetDeleteFaceElementParam DeleteFaceElementParam Getter
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetDeleteFaceElementParam() *FaceElementIdOpenParam {
	return r._deleteFaceElementParam
}

var poolAlibabaDamaiMevOpenDeleteFaceelementAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenDeleteFaceelementRequest()
	},
}

// GetAlibabaDamaiMevOpenDeleteFaceelementRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenDeleteFaceelementAPIRequest
func GetAlibabaDamaiMevOpenDeleteFaceelementAPIRequest() *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest {
	return poolAlibabaDamaiMevOpenDeleteFaceelementAPIRequest.Get().(*AlibabaDamaiMevOpenDeleteFaceelementAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenDeleteFaceelementAPIRequest 将 AlibabaDamaiMevOpenDeleteFaceelementAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeleteFaceelementAPIRequest(v *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeleteFaceelementAPIRequest.Put(v)
}
