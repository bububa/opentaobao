package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取留言列表 APIResponse
taobao.ailab.aicloud.top.message.list

根据指定参数获取留言列表
*/
type TaobaoAilabAicloudTopMessageListAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMessageListResponse
}

type TaobaoAilabAicloudTopMessageListResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_message_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
