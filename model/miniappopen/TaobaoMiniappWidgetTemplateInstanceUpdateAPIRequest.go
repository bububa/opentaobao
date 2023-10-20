package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappwidgettemplateinstanceupdateAPIRequest 小部件实例化版本更新 API请求
// taobao.miniapp.widget.template.instance.update
//
// 小部件版本更新
type TaobaominiappwidgettemplateinstanceupdateAPIRequest struct {
	model.Params
	// 参数信息
	_paramMiniAppInstantiateTemplateAppUpdateRequest *MiniAppInstantiateTemplateAppUpdateRequest
}

// NewTaobaominiappwidgettemplateinstanceupdateRequest 初始化TaobaominiappwidgettemplateinstanceupdateAPIRequest对象
func NewTaobaominiappwidgettemplateinstanceupdateRequest() *TaobaominiappwidgettemplateinstanceupdateAPIRequest {
	return &TaobaominiappwidgettemplateinstanceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappwidgettemplateinstanceupdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.widget.template.instance.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappwidgettemplateinstanceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappwidgettemplateinstanceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamMiniAppInstantiateTemplateAppUpdateRequest is ParamMiniAppInstantiateTemplateAppUpdateRequest Setter
// 参数信息
func (r *TaobaominiappwidgettemplateinstanceupdateAPIRequest) SetParamMiniAppInstantiateTemplateAppUpdateRequest(_paramMiniAppInstantiateTemplateAppUpdateRequest *MiniAppInstantiateTemplateAppUpdateRequest) error {
	r._paramMiniAppInstantiateTemplateAppUpdateRequest = _paramMiniAppInstantiateTemplateAppUpdateRequest
	r.Set("param_mini_app_instantiate_template_app_update_request", _paramMiniAppInstantiateTemplateAppUpdateRequest)
	return nil
}

// GetParamMiniAppInstantiateTemplateAppUpdateRequest ParamMiniAppInstantiateTemplateAppUpdateRequest Getter
func (r TaobaominiappwidgettemplateinstanceupdateAPIRequest) GetParamMiniAppInstantiateTemplateAppUpdateRequest() *MiniAppInstantiateTemplateAppUpdateRequest {
	return r._paramMiniAppInstantiateTemplateAppUpdateRequest
}
