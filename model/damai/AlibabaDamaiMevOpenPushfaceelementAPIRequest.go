package damai

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
