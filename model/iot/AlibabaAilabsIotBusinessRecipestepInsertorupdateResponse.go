package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
插入或更新食谱步骤 APIResponse
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤
*/
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAilabsIotBusinessRecipestepInsertorupdateResponse `json:"alibaba_ailabs_iot_business_recipestep_insertorupdate_response,omitempty"` 
    AlibabaAilabsIotBusinessRecipestepInsertorupdateResponse
}

/* model for simplify = false
type AlibabaAilabsIotBusinessRecipestepInsertorupdateResponse struct {

    // 信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应code
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 返回结果
    
    RetValue   int64 `json:"ret_value,omitempty"`
    

    // 追踪id
    
    TraceId   string `json:"trace_id,omitempty"`
    

}
*/

type AlibabaAilabsIotBusinessRecipestepInsertorupdateResponse struct {

    // 信息
    Message   string `json:"message,omitempty"`

    // 响应code
    RetCode   int64 `json:"ret_code,omitempty"`

    // 返回结果
    RetValue   int64 `json:"ret_value,omitempty"`

    // 追踪id
    TraceId   string `json:"trace_id,omitempty"`

}
