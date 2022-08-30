package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateDeleteAPIRequest 淘宝短信模板删除 API请求
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
type TaobaoJstSmsTemplateDeleteAPIRequest struct {
	model.Params
	// 删除模板的入参
	_deleteSmsTemplateRequest *TopDeleteSmsTemplateRequest
}

// NewTaobaoJstSmsTemplateDeleteRequest 初始化TaobaoJstSmsTemplateDeleteAPIRequest对象
func NewTaobaoJstSmsTemplateDeleteRequest() *TaobaoJstSmsTemplateDeleteAPIRequest {
	return &TaobaoJstSmsTemplateDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeleteSmsTemplateRequest is DeleteSmsTemplateRequest Setter
// 删除模板的入参
func (r *TaobaoJstSmsTemplateDeleteAPIRequest) SetDeleteSmsTemplateRequest(_deleteSmsTemplateRequest *TopDeleteSmsTemplateRequest) error {
	r._deleteSmsTemplateRequest = _deleteSmsTemplateRequest
	r.Set("delete_sms_template_request", _deleteSmsTemplateRequest)
	return nil
}

// GetDeleteSmsTemplateRequest DeleteSmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetDeleteSmsTemplateRequest() *TopDeleteSmsTemplateRequest {
	return r._deleteSmsTemplateRequest
}
