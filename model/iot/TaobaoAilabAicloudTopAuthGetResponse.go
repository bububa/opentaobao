package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
登陆 APIResponse
taobao.ailab.aicloud.top.auth.get

登陆
*/
type TaobaoAilabAicloudTopAuthGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopAuthGetResponse `json:"ailab_aicloud_top_auth_get_response,omitempty"` 
    TaobaoAilabAicloudTopAuthGetResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopAuthGetResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopAuthGetResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
