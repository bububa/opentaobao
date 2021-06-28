package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取食谱详情 APIResponse
alibaba.ailabs.iot.business.recipe.getdetail

获取食谱详情接口，获取ISV自己的食谱详情数据
*/
type AlibabaAilabsIotBusinessRecipeGetdetailAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAilabsIotBusinessRecipeGetdetailResponse `json:"alibaba_ailabs_iot_business_recipe_getdetail_response,omitempty"` 
    AlibabaAilabsIotBusinessRecipeGetdetailResponse
}

/* model for simplify = false
type AlibabaAilabsIotBusinessRecipeGetdetailResponse struct {

    // 返回包装类
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAilabsIotBusinessRecipeGetdetailResponse struct {

    // 返回包装类
    Result   *BaseResult `json:"result,omitempty"`

}
