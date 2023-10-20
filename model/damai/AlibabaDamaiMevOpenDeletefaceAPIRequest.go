package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeletefaceAPIRequest 大麦换验平台-第三方对外开放-票面接口deleteFace API请求
// alibaba.damai.mev.open.deleteface
//
// deleteFace
type AlibabadamaimevopendeletefaceAPIRequest struct {
	model.Params
	// 入参deleteFaceParam
	_deleteFaceParam *TicketFaceIdOpenParam
}

// NewAlibabadamaimevopendeletefaceRequest 初始化AlibabadamaimevopendeletefaceAPIRequest对象
func NewAlibabadamaimevopendeletefaceRequest() *AlibabadamaimevopendeletefaceAPIRequest {
	return &AlibabadamaimevopendeletefaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopendeletefaceAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopendeletefaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopendeletefaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteFaceParam is DeleteFaceParam Setter
// 入参deleteFaceParam
func (r *AlibabadamaimevopendeletefaceAPIRequest) SetDeleteFaceParam(_deleteFaceParam *TicketFaceIdOpenParam) error {
	r._deleteFaceParam = _deleteFaceParam
	r.Set("delete_face_param", _deleteFaceParam)
	return nil
}

// GetDeleteFaceParam DeleteFaceParam Getter
func (r AlibabadamaimevopendeletefaceAPIRequest) GetDeleteFaceParam() *TicketFaceIdOpenParam {
	return r._deleteFaceParam
}
