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
    Response *TaobaoAilabAicloudTopLikeListResponse `json:"taobao_ailab_aicloud_top_like_list_response,omitempty"`
}

type TaobaoAilabAicloudTopLikeListResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
