package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询食谱 APIResponse
alibaba.ailabs.iot.business.recipe.getpage

分页查询食谱数据
*/
type AlibabaAilabsIotBusinessRecipeGetpageAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAilabsIotBusinessRecipeGetpageResponse `json:"alibaba_ailabs_iot_business_recipe_getpage_response,omitempty"` 
    AlibabaAilabsIotBusinessRecipeGetpageResponse
}

/* model for simplify = false
type AlibabaAilabsIotBusinessRecipeGetpageResponse struct {

    // 返回包装类
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAilabsIotBusinessRecipeGetpageResponse struct {

    // 返回包装类
    Result   *BaseResult `json:"result,omitempty"`

}
