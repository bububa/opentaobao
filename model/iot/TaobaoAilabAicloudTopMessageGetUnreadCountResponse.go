package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取未读的消息数量 APIResponse
taobao.ailab.aicloud.top.message.get.unread.count

开放获取未读留言数量的接口
*/
type TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopMessageGetUnreadCountResponse `json:"taobao_ailab_aicloud_top_message_get_unread_count_response,omitempty"`
}

type TaobaoAilabAicloudTopMessageGetUnreadCountResponse struct {

    // 未读留言的数量
    Model   int64 `json:"model,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 异常描述
    MsgInfo   string `json:"msg_info,omitempty"`

}
