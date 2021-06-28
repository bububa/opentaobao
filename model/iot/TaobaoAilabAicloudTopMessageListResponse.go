package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取留言列表 APIResponse
taobao.ailab.aicloud.top.message.list

根据指定参数获取留言列表
*/
type TaobaoAilabAicloudTopMessageListAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopMessageListResponse `json:"ailab_aicloud_top_message_list_response,omitempty"` 
    TaobaoAilabAicloudTopMessageListResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopMessageListResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopMessageListResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
