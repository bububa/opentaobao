package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单条对话流信息 APIResponse
taobao.ailab.aicloud.top.feedlist.delete

删除指定的某一条对话流信息
*/
type TaobaoAilabAicloudTopFeedlistDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_feedlist_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // model
    
    Model   string `json:"model,omitempty" xml:"