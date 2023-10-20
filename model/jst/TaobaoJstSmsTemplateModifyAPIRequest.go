package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatemodifyAPIRequest 淘宝短信模板修改 API请求
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
type TaobaojstsmstemplatemodifyAPIRequest struct {
	model.Params
	// 修改模板的入参
	_modifySmsTemplateRequest *TopModifySmsTemplateRequest
}

// NewTaobaojstsmstemplatemodifyRequest 初始化TaobaojstsmstemplatemodifyAPIRequest对象
func NewTaobaojstsmstemplatemodifyRequest() *TaobaojstsmstemplatemodifyAPIRequest {
	return &TaobaojstsmstemplatemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmstemplatemodifyAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmstemplatemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmstemplatemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifySmsTemplateRequest is ModifySmsTemplateRequest Setter
// 修改模板的入参
func (r *TaobaojstsmstemplatemodifyAPIRequest) SetModifySmsTemplateRequest(_modifySmsTemplateRequest *TopModifySmsTemplateRequest) error {
	r._modifySmsTemplateRequest = _modifySmsTemplateRequest
	r.Set("modify_sms_template_request", _modifySmsTemplateRequest)
	return nil
}

// GetModifySmsTemplateRequest ModifySmsTemplateRequest Getter
func (r TaobaojstsmstemplatemodifyAPIRequest) GetModifySmsTemplateRequest() *TopModifySmsTemplateRequest {
	return r._modifySmsTemplateRequest
}
