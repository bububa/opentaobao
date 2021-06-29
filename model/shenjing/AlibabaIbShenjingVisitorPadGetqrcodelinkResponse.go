package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
pad获取二维码 APIResponse
alibaba.ib.shenjing.visitor.pad.getqrcodelink

pad获取二维码链接。扫码录入人脸。
*/
type AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse struct {
    model.CommonResponse
    AlibabaIbShenjingVisitorPadGetqrcodelinkResponse
}

type AlibabaIbShenjingVisitorPadGetqrcodelinkResponse struct {
    XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_getqrcodelink_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 内容
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`

    
    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 等级
    
    ResultLevel   string `json:"result_level,omitempty" xml:"result_level,omitempty"`

    
    // request_id
    
    ResultRequestId   string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`

    
}
