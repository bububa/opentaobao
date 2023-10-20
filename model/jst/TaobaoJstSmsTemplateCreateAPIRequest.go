package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatecreateAPIRequest 淘宝短信模板创建 API请求
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
type TaobaojstsmstemplatecreateAPIRequest struct {
	model.Params
	// 申请模板入参
	_smsTemplateForIsvRequest *AddSmsTemplateForIsvRequest
}

// NewTaobaojstsmstemplatecreateRequest 初始化TaobaojstsmstemplatecreateAPIRequest对象
func NewTaobaojstsmstemplatecreateRequest() *TaobaojstsmstemplatecreateAPIRequest {
	return &TaobaojstsmstemplatecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmstemplatecreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmstemplatecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmstemplatecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsTemplateForIsvRequest is SmsTemplateForIsvRequest Setter
// 申请模板入参
func (r *TaobaojstsmstemplatecreateAPIRequest) SetSmsTemplateForIsvRequest(_smsTemplateForIsvRequest *AddSmsTemplateForIsvRequest) error {
	r._smsTemplateForIsvRequest = _smsTemplateForIsvRequest
	r.Set("sms_template_for_isv_request", _smsTemplateForIsvRequest)
	return nil
}

// GetSmsTemplateForIsvRequest SmsTemplateForIsvRequest Getter
func (r TaobaojstsmstemplatecreateAPIRequest) GetSmsTemplateForIsvRequest() *AddSmsTemplateForIsvRequest {
	return r._smsTemplateForIsvRequest
}
