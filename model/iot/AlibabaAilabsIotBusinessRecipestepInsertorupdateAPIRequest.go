package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest 插入或更新食谱步骤 API请求
// alibaba.ailabs.iot.business.recipestep.insertorupdate
//
// 插入或更新食谱步骤
type AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest struct {
	model.Params
	// 食谱步骤开放参数
	_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam
}

// NewAlibabaailabsiotbusinessrecipestepinsertorupdateRequest 初始化AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest对象
func NewAlibabaailabsiotbusinessrecipestepinsertorupdateRequest() *AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest {
	return &AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipestep.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusinessRecipeStepOpenParam is ParamBusinessRecipeStepOpenParam Setter
// 食谱步骤开放参数
func (r *AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest) SetParamBusinessRecipeStepOpenParam(_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam) error {
	r._paramBusinessRecipeStepOpenParam = _paramBusinessRecipeStepOpenParam
	r.Set("param_business_recipe_step_open_param", _paramBusinessRecipeStepOpenParam)
	return nil
}

// GetParamBusinessRecipeStepOpenParam ParamBusinessRecipeStepOpenParam Getter
func (r AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest) GetParamBusinessRecipeStepOpenParam() *BusinessRecipeStepOpenParam {
	return r._paramBusinessRecipeStepOpenParam
}
