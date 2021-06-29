package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
过滤列表歌曲存在于收藏列表的 APIResponse
taobao.ailab.aicloud.top.like.filter

过滤出传入列表歌曲存在于收藏列表的
*/
type TaobaoAilabAicloudTopLikeFilterAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopLikeFilterResponse
}

type TaobaoAilabAicloudTopLikeFilterResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_like_filter_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
