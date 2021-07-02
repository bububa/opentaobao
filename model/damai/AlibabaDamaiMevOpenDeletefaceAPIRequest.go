package damai

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletefaceAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletefaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
