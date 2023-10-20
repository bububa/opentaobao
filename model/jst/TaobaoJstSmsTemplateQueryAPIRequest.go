package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatequeryAPIRequest 淘宝短信模板查询 API请求
// taobao.jst.sms.template.query
//
// 淘宝短信模板查询
type TaobaojstsmstemplatequeryAPIRequest struct {
	model.Params
	// 查询短信模板的入参
	_querySmsTemplateRequest *TopQuerySmsTemplateRequest
}

// NewTaobaojstsmstemplatequeryRequest 初始化TaobaojstsmstemplatequeryAPIRequest对象
func NewTaobaojstsmstemplatequeryRequest() *TaobaojstsmstemplatequeryAPIRequest {
	return &TaobaojstsmstemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmstemplatequeryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmstemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmstemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuerySmsTemplateRequest is QuerySmsTemplateRequest Setter
// 查询短信模板的入参
func (r *TaobaojstsmstemplatequeryAPIRequest) SetQuerySmsTemplateRequest(_querySmsTemplateRequest *TopQuerySmsTemplateRequest) error {
	r._querySmsTemplateRequest = _querySmsTemplateRequest
	r.Set("query_sms_template_request", _querySmsTemplateRequest)
	return nil
}

// GetQuerySmsTemplateRequest QuerySmsTemplateRequest Getter
func (r TaobaojstsmstemplatequeryAPIRequest) GetQuerySmsTemplateRequest() *TopQuerySmsTemplateRequest {
	return r._querySmsTemplateRequest
}
