package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
插入或更新食谱步骤 APIRequest
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤
*/
type AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest struct {
    model.Params

    // 食谱步骤开放参数
    paramBusinessRecipeStepOpenParam   *BusinessRecipeStepOpenParam 

}

func NewAlibabaAilabsIotBusinessRecipestepInsertorupdateRequest() *AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest{
    return &AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipestep.insertorupdate"
}

func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest) SetParamBusinessRecipeStepOpenParam(paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam) error {
    r.paramBusinessRecipeStepOpenParam = paramBusinessRecipeStepOpenParam
    r.Set("param_business_recipe_step_open_param", paramBusinessRecipeStepOpenParam)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest) GetParamBusinessRecipeStepOpenParam() *BusinessRecipeStepOpenParam {
    return r.paramBusinessRecipeStepOpenParam
}

