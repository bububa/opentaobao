package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
解除码的关联关系 APIResponse
alibaba.alihealth.tracecodeseller.code.relation.codeantiactive

解除码的关联关系
*/
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveResponse
}

type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_code_relation_codeantiactive_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功失败标记
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 成功失败信息编码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 成功失败信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
