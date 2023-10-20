package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest 插入和更新食谱 API请求
// alibaba.ailabs.iot.business.recipe.insertorupdate
//
// 插入和更新食谱，将isv的食谱添加到云端进行存储
type AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest struct {
	model.Params
	// 行业食谱开放参数
	_paramBusinessRecipeOpenParam *BusinessRecipeOpenParam
}

// NewAlibabaailabsiotbusinessrecipeinsertorupdateRequest 初始化AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest对象
func NewAlibabaailabsiotbusinessrecipeinsertorupdateRequest() *AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest {
	return &AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusinessRecipeOpenParam is ParamBusinessRecipeOpenParam Setter
// 行业食谱开放参数
func (r *AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest) SetParamBusinessRecipeOpenParam(_paramBusinessRecipeOpenParam *BusinessRecipeOpenParam) error {
	r._paramBusinessRecipeOpenParam = _paramBusinessRecipeOpenParam
	r.Set("param_business_recipe_open_param", _paramBusinessRecipeOpenParam)
	return nil
}

// GetParamBusinessRecipeOpenParam ParamBusinessRecipeOpenParam Getter
func (r AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest) GetParamBusinessRecipeOpenParam() *BusinessRecipeOpenParam {
	return r._paramBusinessRecipeOpenParam
}
