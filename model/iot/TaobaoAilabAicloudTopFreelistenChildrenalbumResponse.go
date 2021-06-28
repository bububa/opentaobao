package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
儿童音频列表 APIResponse
taobao.ailab.aicloud.top.freelisten.childrenalbum

儿童音频列表
*/
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopFreelistenChildrenalbumResponse `json:"ailab_aicloud_top_freelisten_childrenalbum_response,omitempty"` 
    TaobaoAilabAicloudTopFreelistenChildrenalbumResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopFreelistenChildrenalbumResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopFreelistenChildrenalbumResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
