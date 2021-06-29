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
    TaobaoAilabAicloudTopLikeDeleteResponse
}

type TaobaoAilabAicloudTopLikeDeleteResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_like_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 具体结果值
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
