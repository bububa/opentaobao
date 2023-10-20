package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushfaceAPIRequest 大麦换验平台-第三方对外开放-票面接口pushFace API请求
// alibaba.damai.mev.open.pushface
//
// pushFace
type AlibabadamaimevopenpushfaceAPIRequest struct {
	model.Params
	// 入参pushFaceParam
	_pushFaceParam *ThirdTicketFacePushOpenParam
}

// NewAlibabadamaimevopenpushfaceRequest 初始化AlibabadamaimevopenpushfaceAPIRequest对象
func NewAlibabadamaimevopenpushfaceRequest() *AlibabadamaimevopenpushfaceAPIRequest {
	return &AlibabadamaimevopenpushfaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushfaceAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushfaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushfaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushFaceParam is PushFaceParam Setter
// 入参pushFaceParam
func (r *AlibabadamaimevopenpushfaceAPIRequest) SetPushFaceParam(_pushFaceParam *ThirdTicketFacePushOpenParam) error {
	r._pushFaceParam = _pushFaceParam
	r.Set("push_face_param", _pushFaceParam)
	return nil
}

// GetPushFaceParam PushFaceParam Getter
func (r AlibabadamaimevopenpushfaceAPIRequest) GetPushFaceParam() *ThirdTicketFacePushOpenParam {
	return r._pushFaceParam
}
