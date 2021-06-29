package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
机翻Api授权 APIResponse
alibaba.seaking.authmachineapi

机翻Api授权
*/
type AlibabaSeakingAuthmachineapiAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingAuthmachineapiResponse
}

type AlibabaSeakingAuthmachineapiResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_authmachineapi_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
