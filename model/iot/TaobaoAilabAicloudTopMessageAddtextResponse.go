package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵代说 API返回值 
taobao.ailab.aicloud.top.message.addtext

精灵代说
*/
type TaobaoAilabAicloudTopMessageAddtextAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMessageAddtextResponse
}

// 精灵代说 成功返回结果
type TaobaoAilabAicloudTopMessageAddtextResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_message_addtext_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
