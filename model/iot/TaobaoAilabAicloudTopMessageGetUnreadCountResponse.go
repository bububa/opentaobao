package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取未读的消息数量 APIResponse
taobao.ailab.aicloud.top.message.get.unread.count

开放获取未读留言数量的接口
*/
type TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_message_get_unread_count_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 未读留言的数量
    
    Model   int64 `json:"model,omitempty" xml:"