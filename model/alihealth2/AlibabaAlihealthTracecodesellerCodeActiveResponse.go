package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传激活码的文件 APIResponse
alibaba.alihealth.tracecodeseller.code.active

上传商品的激活码文件，存到系统中
*/
type AlibabaAlihealthTracecodesellerCodeActiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerCodeActiveResponse
}

type AlibabaAlihealthTracecodesellerCodeActiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_code_active_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
