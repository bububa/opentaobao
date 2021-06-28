package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
音频相关鉴权接口 APIResponse
alibaba.interact.media.audio

新音频包的鉴权接口
*/
type AlibabaInteractMediaAudioAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_media_audio_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

}
