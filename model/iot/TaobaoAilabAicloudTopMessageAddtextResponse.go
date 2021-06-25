package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
精灵代说 APIResponse
taobao.ailab.aicloud.top.message.addtext

精灵代说
*/
type TaobaoAilabAicloudTopMessageAddtextAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopMessageAddtextResponse `json:"taobao_ailab_aicloud_top_message_addtext_response,omitempty"`
}

type TaobaoAilabAicloudTopMessageAddtextResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
