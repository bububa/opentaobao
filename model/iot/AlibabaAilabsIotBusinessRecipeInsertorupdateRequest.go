package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
插入和更新食谱 APIRequest
alibaba.ailabs.iot.business.recipe.insertorupdate

插入和更新食谱，将isv的食谱添加到云端进行存储
*/
type AlibabaAilabsIotBusinessRecipeInsertorupdateRequest struct {
    model.Params

    // 行业食谱开放参数
    paramBusinessRecipeOpenParam   *BusinessRecipeOpenParam 

}

func NewAlibabaAilabsIotBusinessRecipeInsertorupdateRequest() *AlibabaAilabsIotBusinessRecipeInsertorupdateRequest{
    return &AlibabaAilabsIotBusinessRecipeInsertorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.insertorupdate"
}

func (r AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) SetParamBusinessRecipeOpenParam(paramBusinessRecipeOpenParam *BusinessRecipeOpenParam) error {
    r.paramBusinessRecipeOpenParam = paramBusinessRecipeOpenParam
    r.Set("param_business_recipe_open_param", paramBusinessRecipeOpenParam)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipeInsertorupdateRequest) GetParamBusinessRecipeOpenParam() *BusinessRecipeOpenParam {
    return r.paramBusinessRecipeOpenParam
}

