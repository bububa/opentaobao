package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
过滤列表歌曲存在于收藏列表的 API返回值 
taobao.ailab.aicloud.top.like.filter

过滤出传入列表歌曲存在于收藏列表的
*/
type TaobaoAilabAicloudTopLikeFilterAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopLikeFilterAPIResponseModel
}

// 过滤列表歌曲存在于收藏列表的 成功返回结果
type TaobaoAilabAicloudTopLikeFilterAPIResponseModel struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_like_filter_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
