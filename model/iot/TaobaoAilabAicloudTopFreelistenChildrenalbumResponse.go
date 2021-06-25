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
    Response *TaobaoAilabAicloudTopFreelistenChildrenalbumResponse `json:"taobao_ailab_aicloud_top_freelisten_childrenalbum_response,omitempty"`
}

type TaobaoAilabAicloudTopFreelistenChildrenalbumResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
