package dmp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodmpcrowdtemplateapplyAPIRequest 人群模版采纳并生成人群API API请求
// taobao.dmp.crowd.template.apply
//
// 人群模版采纳并生成人群API
type TaobaodmpcrowdtemplateapplyAPIRequest struct {
	model.Params
	// 请求体
	_apiContext *ApiContextDto
	// 采纳模版context
	_templateContext *TemplateContextDto
}

// NewTaobaodmpcrowdtemplateapplyRequest 初始化TaobaodmpcrowdtemplateapplyAPIRequest对象
func NewTaobaodmpcrowdtemplateapplyRequest() *TaobaodmpcrowdtemplateapplyAPIRequest {
	return &TaobaodmpcrowdtemplateapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodmpcrowdtemplateapplyAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.template.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodmpcrowdtemplateapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodmpcrowdtemplateapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiContext is ApiContext Setter
// 请求体
func (r *TaobaodmpcrowdtemplateapplyAPIRequest) SetApiContext(_apiContext *ApiContextDto) error {
	r._apiContext = _apiContext
	r.Set("api_context", _apiContext)
	return nil
}

// GetApiContext ApiContext Getter
func (r TaobaodmpcrowdtemplateapplyAPIRequest) GetApiContext() *ApiContextDto {
	return r._apiContext
}

// SetTemplateContext is TemplateContext Setter
// 采纳模版context
func (r *TaobaodmpcrowdtemplateapplyAPIRequest) SetTemplateContext(_templateContext *TemplateContextDto) error {
	r._templateContext = _templateContext
	r.Set("template_context", _templateContext)
	return nil
}

// GetTemplateContext TemplateContext Getter
func (r TaobaodmpcrowdtemplateapplyAPIRequest) GetTemplateContext() *TemplateContextDto {
	return r._templateContext
}
