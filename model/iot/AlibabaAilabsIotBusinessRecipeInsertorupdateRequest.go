package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
插入和更新食谱 API请求
alibaba.ailabs.iot.business.recipe.insertorupdate

插入和更新食谱，将isv的食谱添加到云端进行存储
*/
type AlibabaAilabsIotBusinessRecipeInsertorupdateRequest struct {
    model.Params
    // 行业食谱开放参数
    _paramBusinessRecipeOpenParam   *BusinessRecipeOpenParam
}

// 初始化AlibabaAilabsIotBusinessRecipeInsertorupdateRequest对象
func NewAlibabaAilabsIotBusinessRecipeInsertorupdateRequest() *AlibabaAilabsIotBusinessRecipeInsertorupdateRequest{
    return &AlibabaAilabsIotBusinessRecipeInsertorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.insertorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBusinessRecipeOpenParam Setter
// 行业食谱开放参数
func (r *AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) SetParamBusinessRecipeOpenParam(_paramBusinessRecipeOpenParam *BusinessRecipeOpenParam) error {
    r._paramBusinessRecipeOpenParam = _paramBusinessRecipeOpenParam
    r.Set("param_business_recipe_open_param", _paramBusinessRecipeOpenParam)
    return nil
}

// ParamBusinessRecipeOpenParam Getter
func (r AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) GetParamBusinessRecipeOpenParam() *BusinessRecipeOpenParam {
    return r._paramBusinessRecipeOpenParam
}
