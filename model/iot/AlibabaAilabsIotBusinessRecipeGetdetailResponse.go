package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取食谱详情 APIResponse
alibaba.ailabs.iot.business.recipe.getdetail

获取食谱详情接口，获取ISV自己的食谱详情数据
*/
type AlibabaAilabsIotBusinessRecipeGetdetailAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotBusinessRecipeGetdetailResponse
}

type AlibabaAilabsIotBusinessRecipeGetdetailResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipe_getdetail_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
