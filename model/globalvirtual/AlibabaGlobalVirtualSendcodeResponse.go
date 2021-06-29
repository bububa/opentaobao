package globalvirtual

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际虚拟商品发码服务 APIResponse
alibaba.global.virtual.sendcode

global virtual send code service
*/
type AlibabaGlobalVirtualSendcodeAPIResponse struct {
    model.CommonResponse
    AlibabaGlobalVirtualSendcodeResponse
}

type AlibabaGlobalVirtualSendcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_global_virtual_sendcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result describe
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
