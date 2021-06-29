package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
故事机发送文本留言 APIResponse
taobao.ailab.aicloud.top.message.sendtext

故事机文本留言
*/
type TaobaoAilabAicloudTopMessageSendtextAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMessageSendtextResponse
}

type TaobaoAilabAicloudTopMessageSendtextResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_message_sendtext_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
