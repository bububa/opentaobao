package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传激活码的文件 API返回值 
alibaba.alihealth.tracecodeseller.code.active

上传商品的激活码文件，存到系统中
*/
type AlibabaAlihealthTracecodesellerCodeActiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerCodeActiveResponse
}

// 上传激活码的文件 成功返回结果
type AlibabaAlihealthTracecodesellerCodeActiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_code_active_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
