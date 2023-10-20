package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQncopilotPictureAuditAPIRequest AIGC创作图片审核 API请求
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
type TaobaoQncopilotPictureAuditAPIRequest struct {
	model.Params
	// 图片审核请求
	_param *PicAuditParam
}

// NewTaobaoQncopilotPictureAuditRequest 初始化TaobaoQncopilotPictureAuditAPIRequest对象
func NewTaobaoQncopilotPictureAuditRequest() *TaobaoQncopilotPictureAuditAPIRequest {
	return &TaobaoQncopilotPictureAuditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQncopilotPictureAuditAPIRequest) GetApiMethodName() string {
	return "taobao.qncopilot.picture.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQncopilotPictureAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQncopilotPictureAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 图片审核请求
func (r *TaobaoQncopilotPictureAuditAPIRequest) SetParam(_param *PicAuditParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoQncopilotPictureAuditAPIRequest) GetParam() *PicAuditParam {
	return r._param
}
