package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
插入和更新食谱 APIResponse
alibaba.ailabs.iot.business.recipe.insertorupdate

插入和更新食谱，将isv的食谱添加到云端进行存储
*/
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotBusinessRecipeInsertorupdateResponse
}

type AlibabaAilabsIotBusinessRecipeInsertorupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipe_insertorupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应code
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 返回结果，行业食谱Id
    
    RetValue   int64 `json:"ret_value,omitempty" xml:"ret_value,omitempty"`

    
    // 追踪id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
}
