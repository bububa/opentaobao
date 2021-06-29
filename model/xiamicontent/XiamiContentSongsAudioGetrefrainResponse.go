package xiamicontent

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取副歌信息 API返回值 
xiami.content.songs.audio.getrefrain

获取歌曲音频副歌
*/
type XiamiContentSongsAudioGetrefrainAPIResponse struct {
    model.CommonResponse
    XiamiContentSongsAudioGetrefrainResponse
}

// 获取副歌信息 成功返回结果
type XiamiContentSongsAudioGetrefrainResponse struct {
    XMLName xml.Name `xml:"xiami_content_songs_audio_getrefrain_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求结果信息
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 音频信息
    Audios   []SongAudiosDTO `json:"audios,omitempty" xml:"audios>song_audios_dto,omitempty"`
}
