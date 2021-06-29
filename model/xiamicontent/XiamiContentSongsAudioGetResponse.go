package xiamicontent

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲音频 APIResponse
xiami.content.songs.audio.get

获取歌曲音频
*/
type XiamiContentSongsAudioGetAPIResponse struct {
    model.CommonResponse
    XiamiContentSongsAudioGetResponse
}

type XiamiContentSongsAudioGetResponse struct {
    XMLName xml.Name `xml:"xiami_content_songs_audio_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果信息
    
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 音频信息
    
    Audios   []SongAudiosDto `json:"audios,omitempty" xml:"audios>song_audios_dto,omitempty"`
    
    
}
