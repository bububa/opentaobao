package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatereportAPIRequest 服务商存量短信模板上传 API请求
// taobao.jst.sms.template.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的模板信息，确保模板数据正确，否则会导致短信发送失败！！！
type TaobaojstsmstemplatereportAPIRequest struct {
	model.Params
	// 存量短信模板上报入参
	_smsTemplateRequest *SmsTemplateRequest
}

// NewTaobaojstsmstemplatereportRequest 初始化TaobaojstsmstemplatereportAPIRequest对象
func NewTaobaojstsmstemplatereportRequest() *TaobaojstsmstemplatereportAPIRequest {
	return &TaobaojstsmstemplatereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmstemplatereportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmstemplatereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmstemplatereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsTemplateRequest is SmsTemplateRequest Setter
// 存量短信模板上报入参
func (r *TaobaojstsmstemplatereportAPIRequest) SetSmsTemplateRequest(_smsTemplateRequest *SmsTemplateRequest) error {
	r._smsTemplateRequest = _smsTemplateRequest
	r.Set("sms_template_request", _smsTemplateRequest)
	return nil
}

// GetSmsTemplateRequest SmsTemplateRequest Getter
func (r TaobaojstsmstemplatereportAPIRequest) GetSmsTemplateRequest() *SmsTemplateRequest {
	return r._smsTemplateRequest
}
