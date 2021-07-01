package alimsg

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
虚拟发货消息发送接口 API返回值 
alibaba.idle.order.msg.send

用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
*/
type AlibabaIdleOrderMsgSendAPIResponse struct {
    model.CommonResponse
    AlibabaIdleOrderMsgSendAPIResponseModel
}

// 虚拟发货消息发送接口 成功返回结果
type AlibabaIdleOrderMsgSendAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_idle_order_msg_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否发送成功
    SendSuccess   bool `json:"send_success,omitempty" xml:"send_success,omitempty"`
}
