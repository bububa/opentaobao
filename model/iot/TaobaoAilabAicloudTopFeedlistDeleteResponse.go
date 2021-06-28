package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除单条对话流信息 APIResponse
taobao.ailab.aicloud.top.feedlist.delete

删除指定的某一条对话流信息
*/
type TaobaoAilabAicloudTopFeedlistDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopFeedlistDeleteResponse `json:"ailab_aicloud_top_feedlist_delete_response,omitempty"` 
    TaobaoAilabAicloudTopFeedlistDeleteResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopFeedlistDeleteResponse struct {

    // model
    
    Model   string `json:"model,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    IsSuccess   string `json:"is_success,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopFeedlistDeleteResponse struct {

    // model
    Model   string `json:"model,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    IsSuccess   string `json:"is_success,omitempty"`

}
