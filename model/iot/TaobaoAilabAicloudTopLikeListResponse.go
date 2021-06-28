package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
列出收藏列表 APIResponse
taobao.ailab.aicloud.top.like.list

列出收藏列表
*/
type TaobaoAilabAicloudTopLikeListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_like_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"