package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
过滤列表歌曲存在于收藏列表的 APIResponse
taobao.ailab.aicloud.top.like.filter

过滤出传入列表歌曲存在于收藏列表的
*/
type TaobaoAilabAicloudTopLikeFilterAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopLikeFilterResponse `json:"taobao_ailab_aicloud_top_like_filter_response,omitempty"`
}

type TaobaoAilabAicloudTopLikeFilterResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
