package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
插入或更新食谱步骤 APIResponse
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤
*/
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotBusinessRecipestepInsertorupdateResponse
}

type AlibabaAilabsIotBusinessRecipestepInsertorupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipestep_insertorupdate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应code
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 返回结果
    
    RetValue   int64 `json:"ret_value,omitempty" xml:"ret_value,omitempty"`

    
    // 追踪id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
}
