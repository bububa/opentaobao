package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 APIResponse
alibaba.interact.ui.video

Weex页面播放视频
*/
type AlibabaInteractUiVideoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_ui_video_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 仅作鉴权使用，没有实际数据传输
    
    Unnamed   string `json:"unnamed,omitempty" xml:"