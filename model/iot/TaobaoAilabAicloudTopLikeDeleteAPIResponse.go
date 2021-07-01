package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消收藏 API返回值 
taobao.ailab.aicloud.top.like.delete

取消收藏
*/
type TaobaoAilabAicloudTopLikeDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopLikeDeleteAPIResponseModel
}

// 取消收藏 成功返回结果
type TaobaoAilabAicloudTopLikeDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_like_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 具体结果值
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}
