package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加收藏 APIResponse
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏
*/
type TaobaoAilabAicloudTopLikeAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_like_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 具体信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"