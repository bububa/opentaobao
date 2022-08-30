package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateCreateAPIRequest 淘宝短信模板创建 API请求
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
type TaobaoJstSmsTemplateCreateAPIRequest struct {
	model.Params
	// 申请模板入参
	_smsTemplateForIsvRequest *AddSmsTemplateForIsvRequest
}

// NewTaobaoJstSmsTemplateCreateRequest 初始化TaobaoJstSmsTemplateCreateAPIRequest对象
func NewTaobaoJstSmsTemplateCreateRequest() *TaobaoJstSmsTemplateCreateAPIRequest {
	return &TaobaoJstSmsTemplateCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSmsTemplateForIsvRequest is SmsTemplateForIsvRequest Setter
// 申请模板入参
func (r *TaobaoJstSmsTemplateCreateAPIRequest) SetSmsTemplateForIsvRequest(_smsTemplateForIsvRequest *AddSmsTemplateForIsvRequest) error {
	r._smsTemplateForIsvRequest = _smsTemplateForIsvRequest
	r.Set("sms_template_for_isv_request", _smsTemplateForIsvRequest)
	return nil
}

// GetSmsTemplateForIsvRequest SmsTemplateForIsvRequest Getter
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetSmsTemplateForIsvRequest() *AddSmsTemplateForIsvRequest {
	return r._smsTemplateForIsvRequest
}
