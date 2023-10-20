package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest 插入或更新食谱步骤 API请求
// alibaba.ailabs.iot.business.recipestep.insertorupdate
//
// 插入或更新食谱步骤
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest struct {
	model.Params
	// 食谱步骤开放参数
	_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam
}

// NewAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest 初始化AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest对象
func NewAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest() *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest {
	return &AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) Reset() {
	r._paramBusinessRecipeStepOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipestep.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusinessRecipeStepOpenParam is ParamBusinessRecipeStepOpenParam Setter
// 食谱步骤开放参数
func (r *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) SetParamBusinessRecipeStepOpenParam(_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam) error {
	r._paramBusinessRecipeStepOpenParam = _paramBusinessRecipeStepOpenParam
	r.Set("param_business_recipe_step_open_param", _paramBusinessRecipeStepOpenParam)
	return nil
}

// GetParamBusinessRecipeStepOpenParam ParamBusinessRecipeStepOpenParam Getter
func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) GetParamBusinessRecipeStepOpenParam() *BusinessRecipeStepOpenParam {
	return r._paramBusinessRecipeStepOpenParam
}

var poolAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest()
	},
}

// GetAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest 从 sync.Pool 获取 AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest
func GetAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest() *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest {
	return poolAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest.Get().(*AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest)
}

// ReleaseAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest 将 AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest(v *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest.Put(v)
}
