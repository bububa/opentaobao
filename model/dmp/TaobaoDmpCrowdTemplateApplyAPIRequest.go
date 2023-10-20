package dmp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDmpCrowdTemplateApplyAPIRequest) Reset() {
	r._apiContext = nil
	r._templateContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.template.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDmpCrowdTemplateApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoDmpCrowdTemplateApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDmpCrowdTemplateApplyRequest()
	},
}

// GetTaobaoDmpCrowdTemplateApplyRequest 从 sync.Pool 获取 TaobaoDmpCrowdTemplateApplyAPIRequest
func GetTaobaoDmpCrowdTemplateApplyAPIRequest() *TaobaoDmpCrowdTemplateApplyAPIRequest {
	return poolTaobaoDmpCrowdTemplateApplyAPIRequest.Get().(*TaobaoDmpCrowdTemplateApplyAPIRequest)
}

// ReleaseTaobaoDmpCrowdTemplateApplyAPIRequest 将 TaobaoDmpCrowdTemplateApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoDmpCrowdTemplateApplyAPIRequest(v *TaobaoDmpCrowdTemplateApplyAPIRequest) {
	v.Reset()
	poolTaobaoDmpCrowdTemplateApplyAPIRequest.Put(v)
}
