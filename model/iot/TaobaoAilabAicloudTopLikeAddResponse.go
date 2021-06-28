package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加收藏 APIResponse
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏
*/
type TaobaoAilabAicloudTopLikeAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopLikeAddResponse `json:"ailab_aicloud_top_like_add_response,omitempty"` 
    TaobaoAilabAicloudTopLikeAddResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopLikeAddResponse struct {

    // 具体信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 标志是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 具体的结果值
    
    Model   bool `json:"model,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopLikeAddResponse struct {

    // 具体信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 标志是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 具体的结果值
    Model   bool `json:"model,omitempty"`

}
