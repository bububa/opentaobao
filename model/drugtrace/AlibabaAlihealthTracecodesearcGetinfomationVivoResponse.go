package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取vivo banner APIResponse
alibaba.alihealth.tracecodesearc.getinfomation.vivo

获取vivo banner  url
*/
type AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesearcGetinfomationVivoResponse
}

type AlibabaAlihealthTracecodesearcGetinfomationVivoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodesearc_getinfomation_vivo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // bannerURL
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 操作说明
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 操作码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
