package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取对话流列表 APIResponse
taobao.ailab.aicloud.top.feedlist.get

获取指定应用的对话流信息
*/
type TaobaoAilabAicloudTopFeedlistGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_feedlist_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"