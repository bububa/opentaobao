package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消收藏 APIResponse
taobao.ailab.aicloud.top.like.delete

取消收藏
*/
type TaobaoAilabAicloudTopLikeDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopLikeDeleteResponse `json:"ailab_aicloud_top_like_delete_response,omitempty"` 
    TaobaoAilabAicloudTopLikeDeleteResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopLikeDeleteResponse struct {

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 具体结果值
    
    Model   bool `json:"model,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopLikeDeleteResponse struct {

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 具体结果值
    Model   bool `json:"model,omitempty"`

}
