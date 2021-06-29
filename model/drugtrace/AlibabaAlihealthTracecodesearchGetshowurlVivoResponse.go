package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取药品扫码落地页vivo APIResponse
alibaba.alihealth.tracecodesearch.getshowurl.vivo

获取药品扫码落地页vivo
*/
type AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesearchGetshowurlVivoResponse
}

type AlibabaAlihealthTracecodesearchGetshowurlVivoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodesearch_getshowurl_vivo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 落地页地址
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 操作说明
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 操作码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
