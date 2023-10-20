package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappWidgetTemplateInstantiateAPIRequest 小部件实例化接口 API请求
// taobao.miniapp.widget.template.instantiate
//
// 小部件实例化接口
type TaobaoMiniappWidgetTemplateInstantiateAPIRequest struct {
	model.Params
	// 实例化数据
	_paramMiniAppInstantiateTemplateAppSimpleRequest *MiniAppInstantiateTemplateAppSimpleRequest
}

// NewTaobaoMiniappWidgetTemplateInstantiateRequest 初始化TaobaoMiniappWidgetTemplateInstantiateAPIRequest对象
func NewTaobaoMiniappWidgetTemplateInstantiateRequest() *TaobaoMiniappWidgetTemplateInstantiateAPIRequest {
	return &TaobaoMiniappWidgetTemplateInstantiateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappWidgetTemplateInstantiateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.widget.template.instantiate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappWidgetTemplateInstantiateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappWidgetTemplateInstantiateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamMiniAppInstantiateTemplateAppSimpleRequest is ParamMiniAppInstantiateTemplateAppSimpleRequest Setter
// 实例化数据
func (r *TaobaoMiniappWidgetTemplateInstantiateAPIRequest) SetParamMiniAppInstantiateTemplateAppSimpleRequest(_paramMiniAppInstantiateTemplateAppSimpleRequest *MiniAppInstantiateTemplateAppSimpleRequest) error {
	r._paramMiniAppInstantiateTemplateAppSimpleRequest = _paramMiniAppInstantiateTemplateAppSimpleRequest
	r.Set("param_mini_app_instantiate_template_app_simple_request", _paramMiniAppInstantiateTemplateAppSimpleRequest)
	return nil
}

// GetParamMiniAppInstantiateTemplateAppSimpleRequest ParamMiniAppInstantiateTemplateAppSimpleRequest Getter
func (r TaobaoMiniappWidgetTemplateInstantiateAPIRequest) GetParamMiniAppInstantiateTemplateAppSimpleRequest() *MiniAppInstantiateTemplateAppSimpleRequest {
	return r._paramMiniAppInstantiateTemplateAppSimpleRequest
}
