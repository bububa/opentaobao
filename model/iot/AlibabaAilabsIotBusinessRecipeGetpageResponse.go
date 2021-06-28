package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询食谱 APIResponse
alibaba.ailabs.iot.business.recipe.getpage

分页查询食谱数据
*/
type AlibabaAilabsIotBusinessRecipeGetpageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ailabs_iot_business_recipe_getpage_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"