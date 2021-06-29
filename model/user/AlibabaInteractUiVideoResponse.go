package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 API返回值 
alibaba.interact.ui.video

Weex页面播放视频
*/
type AlibabaInteractUiVideoAPIResponse struct {
    model.CommonResponse
    AlibabaInteractUiVideoResponse
}

// 视频播放 成功返回结果
type AlibabaInteractUiVideoResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_ui_video_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 仅作鉴权使用，没有实际数据传输
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}
