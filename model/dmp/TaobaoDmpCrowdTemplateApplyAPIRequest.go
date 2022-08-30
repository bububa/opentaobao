package dmp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDmpCrowdTemplateApplyAPIRequest 人群模版采纳并生成人群API API请求
// taobao.dmp.crowd.template.apply
//
// 人群模版采纳并生成人群API
type TaobaoDmpCrowdTemplateApplyAPIRequest struct {
	model.Params
	// 请求体
	_apiContext *ApiContextDto
	// 采纳模版context
	_templateContext *TemplateContextDto
}

// NewTaobaoDmpCrowdTemplateApplyRequest 初始化TaobaoDmpCrowdTemplateApplyAPIRequest对象
func NewTaobaoDmpCrowdTemplateApplyRequest() *TaobaoDmpCrowdTemplateApplyAPIRequest {
	return &TaobaoDmpCrowdTemplateApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.template.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiContext is ApiContext Setter
// 请求体
func (r *TaobaoDmpCrowdTemplateApplyAPIRequest) SetApiContext(_apiContext *ApiContextDto) error {
	r._apiContext = _apiContext
	r.Set("api_context", _apiContext)
	return nil
}

// GetApiContext ApiContext Getter
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetApiContext() *ApiContextDto {
	return r._apiContext
}

// SetTemplateContext is TemplateContext Setter
// 采纳模版context
func (r *TaobaoDmpCrowdTemplateApplyAPIRequest) SetTemplateContext(_templateContext *TemplateContextDto) error {
	r._templateContext = _templateContext
	r.Set("template_context", _templateContext)
	return nil
}

// GetTemplateContext TemplateContext Getter
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetTemplateContext() *TemplateContextDto {
	return r._templateContext
}
