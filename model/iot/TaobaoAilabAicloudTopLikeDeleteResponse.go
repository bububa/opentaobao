package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消收藏 APIResponse
taobao.ailab.aicloud.top.like.delete

取消收藏
*/
type TaobaoAilabAicloudTopLikeDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_like_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"