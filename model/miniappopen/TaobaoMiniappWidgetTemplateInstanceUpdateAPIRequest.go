package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest 小部件实例化版本更新 API请求
// taobao.miniapp.widget.template.instance.update
//
// 小部件版本更新
type TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest struct {
	model.Params
	// 参数信息
	_paramMiniAppInstantiateTemplateAppUpdateRequest *MiniAppInstantiateTemplateAppUpdateRequest
}

// NewTaobaoMiniappWidgetTemplateInstanceUpdateRequest 初始化TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest对象
func NewTaobaoMiniappWidgetTemplateInstanceUpdateRequest() *TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest {
	return &TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.widget.template.instance.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamMiniAppInstantiateTemplateAppUpdateRequest is ParamMiniAppInstantiateTemplateAppUpdateRequest Setter
// 参数信息
func (r *TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest) SetParamMiniAppInstantiateTemplateAppUpdateRequest(_paramMiniAppInstantiateTemplateAppUpdateRequest *MiniAppInstantiateTemplateAppUpdateRequest) error {
	r._paramMiniAppInstantiateTemplateAppUpdateRequest = _paramMiniAppInstantiateTemplateAppUpdateRequest
	r.Set("param_mini_app_instantiate_template_app_update_request", _paramMiniAppInstantiateTemplateAppUpdateRequest)
	return nil
}

// GetParamMiniAppInstantiateTemplateAppUpdateRequest ParamMiniAppInstantiateTemplateAppUpdateRequest Getter
func (r TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest) GetParamMiniAppInstantiateTemplateAppUpdateRequest() *MiniAppInstantiateTemplateAppUpdateRequest {
	return r._paramMiniAppInstantiateTemplateAppUpdateRequest
}
