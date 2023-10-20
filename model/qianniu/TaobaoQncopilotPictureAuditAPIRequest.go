package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqncopilotpictureauditAPIRequest AIGC创作图片审核 API请求
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
type TaobaoqncopilotpictureauditAPIRequest struct {
	model.Params
	// 图片审核请求
	_param *PicAuditParam
}

// NewTaobaoqncopilotpictureauditRequest 初始化TaobaoqncopilotpictureauditAPIRequest对象
func NewTaobaoqncopilotpictureauditRequest() *TaobaoqncopilotpictureauditAPIRequest {
	return &TaobaoqncopilotpictureauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqncopilotpictureauditAPIRequest) GetApiMethodName() string {
	return "taobao.qncopilot.picture.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqncopilotpictureauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqncopilotpictureauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 图片审核请求
func (r *TaobaoqncopilotpictureauditAPIRequest) SetParam(_param *PicAuditParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoqncopilotpictureauditAPIRequest) GetParam() *PicAuditParam {
	return r._param
}
