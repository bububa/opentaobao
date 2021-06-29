package alimsg

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
虚拟发货消息发送接口 APIResponse
alibaba.idle.order.msg.send

用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
*/
type AlibabaIdleOrderMsgSendAPIResponse struct {
    model.CommonResponse
    AlibabaIdleOrderMsgSendResponse
}

type AlibabaIdleOrderMsgSendResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_order_msg_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否发送成功
    
    SendSuccess   bool `json:"send_success,omitempty" xml:"send_success,omitempty"`

    
}
