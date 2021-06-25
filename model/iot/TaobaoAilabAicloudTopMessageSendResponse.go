package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发送留言 APIResponse
taobao.ailab.aicloud.top.message.send

供准入的外部用户实现发送留言功能，APP端发送，设备端读取
*/
type TaobaoAilabAicloudTopMessageSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopMessageSendResponse `json:"taobao_ailab_aicloud_top_message_send_response,omitempty"`
}

type TaobaoAilabAicloudTopMessageSendResponse struct {

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // model
    Model   bool `json:"model,omitempty"`

}
