package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatedeleteAPIRequest 淘宝短信模板删除 API请求
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
type TaobaojstsmstemplatedeleteAPIRequest struct {
	model.Params
	// 删除模板的入参
	_deleteSmsTemplateRequest *TopDeleteSmsTemplateRequest
}

// NewTaobaojstsmstemplatedeleteRequest 初始化TaobaojstsmstemplatedeleteAPIRequest对象
func NewTaobaojstsmstemplatedeleteRequest() *TaobaojstsmstemplatedeleteAPIRequest {
	return &TaobaojstsmstemplatedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmstemplatedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmstemplatedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmstemplatedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteSmsTemplateRequest is DeleteSmsTemplateRequest Setter
// 删除模板的入参
func (r *TaobaojstsmstemplatedeleteAPIRequest) SetDeleteSmsTemplateRequest(_deleteSmsTemplateRequest *TopDeleteSmsTemplateRequest) error {
	r._deleteSmsTemplateRequest = _deleteSmsTemplateRequest
	r.Set("delete_sms_template_request", _deleteSmsTemplateRequest)
	return nil
}

// GetDeleteSmsTemplateRequest DeleteSmsTemplateRequest Getter
func (r TaobaojstsmstemplatedeleteAPIRequest) GetDeleteSmsTemplateRequest() *TopDeleteSmsTemplateRequest {
	return r._deleteSmsTemplateRequest
}
