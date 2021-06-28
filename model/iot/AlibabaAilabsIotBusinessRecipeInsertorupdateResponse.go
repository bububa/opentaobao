package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
插入和更新食谱 APIResponse
alibaba.ailabs.iot.business.recipe.insertorupdate

插入和更新食谱，将isv的食谱添加到云端进行存储
*/
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAilabsIotBusinessRecipeInsertorupdateResponse `json:"alibaba_ailabs_iot_business_recipe_insertorupdate_response,omitempty"` 
    AlibabaAilabsIotBusinessRecipeInsertorupdateResponse
}

/* model for simplify = false
type AlibabaAilabsIotBusinessRecipeInsertorupdateResponse struct {

    // 信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应code
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 返回结果，行业食谱Id
    
    RetValue   int64 `json:"ret_value,omitempty"`
    

    // 追踪id
    
    TraceId   string `json:"trace_id,omitempty"`
    

}
*/

type AlibabaAilabsIotBusinessRecipeInsertorupdateResponse struct {

    // 信息
    Message   string `json:"message,omitempty"`

    // 响应code
    RetCode   int64 `json:"ret_code,omitempty"`

    // 返回结果，行业食谱Id
    RetValue   int64 `json:"ret_value,omitempty"`

    // 追踪id
    TraceId   string `json:"trace_id,omitempty"`

}
