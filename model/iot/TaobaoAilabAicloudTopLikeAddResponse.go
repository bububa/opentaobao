package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加收藏 API返回值 
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏
*/
type TaobaoAilabAicloudTopLikeAddAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopLikeAddResponse
}

// 增加收藏 成功返回结果
type TaobaoAilabAicloudTopLikeAddResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_like_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 具体信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 标志是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 具体的结果值
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}
