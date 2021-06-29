package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
本地生活触达模板消息发送接口 APIResponse
alibaba.alsc.chuda.template.send

允许三方小程序通过该api发送本地生活触达消息
*/
type AlibabaAlscChudaTemplateSendAPIResponse struct {
    model.CommonResponse
    AlibabaAlscChudaTemplateSendResponse
}

type AlibabaAlscChudaTemplateSendResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_chuda_template_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgId
    
    ResultObj   int64 `json:"result_obj,omitempty" xml:"result_obj,omitempty"`

    
    // 发送是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 结果码，200表示成功
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 描述信息
    
    ResultDesc   string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`

    
}
