package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateModifyAPIRequest 淘宝短信模板修改 API请求
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
type TaobaoJstSmsTemplateModifyAPIRequest struct {
	model.Params
	// 修改模板的入参
	_modifySmsTemplateRequest *TopModifySmsTemplateRequest
}

// NewTaobaoJstSmsTemplateModifyRequest 初始化TaobaoJstSmsTemplateModifyAPIRequest对象
func NewTaobaoJstSmsTemplateModifyRequest() *TaobaoJstSmsTemplateModifyAPIRequest {
	return &TaobaoJstSmsTemplateModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetModifySmsTemplateRequest is ModifySmsTemplateRequest Setter
// 修改模板的入参
func (r *TaobaoJstSmsTemplateModifyAPIRequest) SetModifySmsTemplateRequest(_modifySmsTemplateRequest *TopModifySmsTemplateRequest) error {
	r._modifySmsTemplateRequest = _modifySmsTemplateRequest
	r.Set("modify_sms_template_request", _modifySmsTemplateRequest)
	return nil
}

// GetModifySmsTemplateRequest ModifySmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetModifySmsTemplateRequest() *TopModifySmsTemplateRequest {
	return r._modifySmsTemplateRequest
}
