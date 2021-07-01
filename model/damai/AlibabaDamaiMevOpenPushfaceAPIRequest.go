package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushfaceAPIRequest
大麦换验平台-第三方对外开放-票面接口pushFace API请求
alibaba.damai.mev.open.pushface

pushFace */
type AlibabaDamaiMevOpenPushfaceAPIRequest struct {
	model.Params
	// 入参pushFaceParam
	_pushFaceParam *ThirdTicketFacePushOpenParam
}

// NewAlibabaDamaiMevOpenPushfaceRequest 初始化AlibabaDamaiMevOpenPushfaceAPIRequest对象
func NewAlibabaDamaiMevOpenPushfaceRequest() *AlibabaDamaiMevOpenPushfaceAPIRequest {
	return &AlibabaDamaiMevOpenPushfaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushFaceParam Setter
// 入参pushFaceParam
func (r *AlibabaDamaiMevOpenPushfaceAPIRequest) SetPushFaceParam(_pushFaceParam *ThirdTicketFacePushOpenParam) error {
	r._pushFaceParam = _pushFaceParam
	r.Set("push_face_param", _pushFaceParam)
	return nil
}

// Get PushFaceParam Getter
func (r AlibabaDamaiMevOpenPushfaceAPIRequest) GetPushFaceParam() *ThirdTicketFacePushOpenParam {
	return r._pushFaceParam
}
