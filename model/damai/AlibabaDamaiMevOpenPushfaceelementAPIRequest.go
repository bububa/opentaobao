package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfaceelementAPIRequest 大麦换验平台-第三方对外开放-票面元素接口pushFaceElement API请求
// alibaba.damai.mev.open.pushfaceelement
//
// pushFaceElement
type AlibabaDamaiMevOpenPushfaceelementAPIRequest struct {
	model.Params
	// 入参pushFaceElementParamList
	_pushFaceElementParamList []ThirdFaceElementPushOpenParam
}

// NewAlibabaDamaiMevOpenPushfaceelementRequest 初始化AlibabaDamaiMevOpenPushfaceelementAPIRequest对象
func NewAlibabaDamaiMevOpenPushfaceelementRequest() *AlibabaDamaiMevOpenPushfaceelementAPIRequest {
	return &AlibabaDamaiMevOpenPushfaceelementAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushfaceelementAPIRequest) Reset() {
	r._pushFaceElementParamList = r._pushFaceElementParamList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfaceelementAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushfaceelement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfaceelementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushfaceelementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushFaceElementParamList is PushFaceElementParamList Setter
// 入参pushFaceElementParamList
func (r *AlibabaDamaiMevOpenPushfaceelementAPIRequest) SetPushFaceElementParamList(_pushFaceElementParamList []ThirdFaceElementPushOpenParam) error {
	r._pushFaceElementParamList = _pushFaceElementParamList
	r.Set("push_face_element_param_list", _pushFaceElementParamList)
	return nil
}

// GetPushFaceElementParamList PushFaceElementParamList Getter
func (r AlibabaDamaiMevOpenPushfaceelementAPIRequest) GetPushFaceElementParamList() []ThirdFaceElementPushOpenParam {
	return r._pushFaceElementParamList
}

var poolAlibabaDamaiMevOpenPushfaceelementAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushfaceelementRequest()
	},
}

// GetAlibabaDamaiMevOpenPushfaceelementRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushfaceelementAPIRequest
func GetAlibabaDamaiMevOpenPushfaceelementAPIRequest() *AlibabaDamaiMevOpenPushfaceelementAPIRequest {
	return poolAlibabaDamaiMevOpenPushfaceelementAPIRequest.Get().(*AlibabaDamaiMevOpenPushfaceelementAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushfaceelementAPIRequest 将 AlibabaDamaiMevOpenPushfaceelementAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushfaceelementAPIRequest(v *AlibabaDamaiMevOpenPushfaceelementAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushfaceelementAPIRequest.Put(v)
}
