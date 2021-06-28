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
    // Response *TaobaoAilabAicloudTopMessageAddtextResponse `json:"ailab_aicloud_top_message_addtext_response,omitempty"` 
    TaobaoAilabAicloudTopMessageAddtextResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopMessageAddtextResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopMessageAddtextResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
