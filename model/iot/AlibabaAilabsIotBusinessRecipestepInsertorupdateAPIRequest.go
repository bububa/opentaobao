package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest
插入或更新食谱步骤 API请求
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤 */
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest struct {
	model.Params
	// 食谱步骤开放参数
	_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam
}

// NewAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest 初始化AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest对象
func NewAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest() *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest {
	return &AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipestep.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBusinessRecipeStepOpenParam Setter
// 食谱步骤开放参数
func (r *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) SetParamBusinessRecipeStepOpenParam(_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam) error {
	r._paramBusinessRecipeStepOpenParam = _paramBusinessRecipeStepOpenParam
	r.Set("param_business_recipe_step_open_param", _paramBusinessRecipeStepOpenParam)
	return nil
}

// Get ParamBusinessRecipeStepOpenParam Getter
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetParamBusinessRecipeStepOpenParam() *BusinessRecipeStepOpenParam {
	return r._paramBusinessRecipeStepOpenParam
}
