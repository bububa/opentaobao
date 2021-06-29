package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发送语音留言 APIResponse
taobao.ailab.aicloud.top.message.sendaudio

将语音的二进制byte[]通过TOP接口发送保存
*/
type TaobaoAilabAicloudTopMessageSendaudioAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMessageSendaudioResponse
}

type TaobaoAilabAicloudTopMessageSendaudioResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_message_sendaudio_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
