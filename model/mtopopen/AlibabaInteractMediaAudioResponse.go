package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
音频相关鉴权接口 APIResponse
alibaba.interact.media.audio

新音频包的鉴权接口
*/
type AlibabaInteractMediaAudioAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractMediaAudioResponse `json:"alibaba_interact_media_audio_response,omitempty"`
}

type AlibabaInteractMediaAudioResponse struct {

}
