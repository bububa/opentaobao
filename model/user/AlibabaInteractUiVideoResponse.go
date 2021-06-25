package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 APIResponse
alibaba.interact.ui.video

Weex页面播放视频
*/
type AlibabaInteractUiVideoAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractUiVideoResponse `json:"alibaba_interact_ui_video_response,omitempty"`
}

type AlibabaInteractUiVideoResponse struct {

    // 仅作鉴权使用，没有实际数据传输
    Unnamed   string `json:"unnamed,omitempty"`

}
