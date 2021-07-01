package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
故事机发送文本留言 API返回值 
taobao.ailab.aicloud.top.message.sendtext

故事机文本留言
*/
type TaobaoAilabAicloudTopMessageSendtextAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMessageSendtextAPIResponseModel
}

// 故事机发送文本留言 成功返回结果
type TaobaoAilabAicloudTopMessageSendtextAPIResponseModel struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_message_sendtext_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
