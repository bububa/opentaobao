package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest 插入和更新食谱 API请求
// alibaba.ailabs.iot.business.recipe.insertorupdate
//
// 插入和更新食谱，将isv的食谱添加到云端进行存储
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest struct {
	model.Params
	// 行业食谱开放参数
	_paramBusinessRecipeOpenParam *BusinessRecipeOpenParam
}

// NewAlibabaAilabsIotBusinessRecipeInsertorupdateRequest 初始化AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest对象
func NewAlibabaAilabsIotBusinessRecipeInsertorupdateRequest() *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest {
	return &AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) Reset() {
	r._paramBusinessRecipeOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusinessRecipeOpenParam is ParamBusinessRecipeOpenParam Setter
// 行业食谱开放参数
func (r *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) SetParamBusinessRecipeOpenParam(_paramBusinessRecipeOpenParam *BusinessRecipeOpenParam) error {
	r._paramBusinessRecipeOpenParam = _paramBusinessRecipeOpenParam
	r.Set("param_business_recipe_open_param", _paramBusinessRecipeOpenParam)
	return nil
}

// GetParamBusinessRecipeOpenParam ParamBusinessRecipeOpenParam Getter
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) GetParamBusinessRecipeOpenParam() *BusinessRecipeOpenParam {
	return r._paramBusinessRecipeOpenParam
}

var poolAlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotBusinessRecipeInsertorupdateRequest()
	},
}

// GetAlibabaAilabsIotBusinessRecipeInsertorupdateRequest 从 sync.Pool 获取 AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest
func GetAlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest() *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest {
	return poolAlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest.Get().(*AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest)
}

// ReleaseAlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest 将 AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest(v *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest.Put(v)
}
