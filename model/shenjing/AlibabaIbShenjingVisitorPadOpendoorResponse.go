package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
访客发起开门 APIResponse
alibaba.ib.shenjing.visitor.pad.opendoor

访客PAD端录入完人脸后，可以点击开门按钮开门。
*/
type AlibabaIbShenjingVisitorPadOpendoorAPIResponse struct {
    model.CommonResponse
    AlibabaIbShenjingVisitorPadOpendoorResponse
}

type AlibabaIbShenjingVisitorPadOpendoorResponse struct {
    XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_opendoor_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 业务处理结果
    
    Content   bool `json:"content,omitempty" xml:"content,omitempty"`

    
    // 请求request_id
    
    ResultRequestId   string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`

    
    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
    // 状态码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 状态描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
