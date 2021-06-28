package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
列出收藏列表 APIResponse
taobao.ailab.aicloud.top.like.list

列出收藏列表
*/
type TaobaoAilabAicloudTopLikeListAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopLikeListResponse `json:"ailab_aicloud_top_like_list_response,omitempty"` 
    TaobaoAilabAicloudTopLikeListResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopLikeListResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopLikeListResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
