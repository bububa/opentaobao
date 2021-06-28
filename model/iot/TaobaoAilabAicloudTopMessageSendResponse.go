package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发送留言 APIResponse
taobao.ailab.aicloud.top.message.send

供准入的外部用户实现发送留言功能，APP端发送，设备端读取
*/
type TaobaoAilabAicloudTopMessageSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_message_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"