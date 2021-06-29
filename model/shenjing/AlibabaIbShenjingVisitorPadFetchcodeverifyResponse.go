package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
访客通过PAD提交访客码 APIResponse
alibaba.ib.shenjing.visitor.pad.fetchcodeverify

访客通过PAD提交访客码，录脸进入园区。
*/
type AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse struct {
    model.CommonResponse
    AlibabaIbShenjingVisitorPadFetchcodeverifyResponse
}

type AlibabaIbShenjingVisitorPadFetchcodeverifyResponse struct {
    XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_fetchcodeverify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 内容
    
    Content   *Content `json:"content,omitempty" xml:"content,omitempty"`

    
    // request_id
    
    ResultRequestId   string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`

    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
