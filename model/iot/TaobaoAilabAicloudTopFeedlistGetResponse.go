package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取对话流列表 APIResponse
taobao.ailab.aicloud.top.feedlist.get

获取指定应用的对话流信息
*/
type TaobaoAilabAicloudTopFeedlistGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopFeedlistGetResponse `json:"taobao_ailab_aicloud_top_feedlist_get_response,omitempty"`
}

type TaobaoAilabAicloudTopFeedlistGetResponse struct {

    // 返回值
    Result   *AiCloudResult `json:"result,omitempty"`

}
