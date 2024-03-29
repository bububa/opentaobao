package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsAudioGetrefrainAPIResponse 获取副歌信息 API返回值
// xiami.content.songs.audio.getrefrain
//
// 获取歌曲音频副歌
type XiamiContentSongsAudioGetrefrainAPIResponse struct {
	model.CommonResponse
	XiamiContentSongsAudioGetrefrainAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentSongsAudioGetrefrainAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentSongsAudioGetrefrainAPIResponseModel).Reset()
}

// XiamiContentSongsAudioGetrefrainAPIResponseModel is 获取副歌信息 成功返回结果
type XiamiContentSongsAudioGetrefrainAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_audio_getrefrain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 音频信息
	Audios []SongAudiosDto `json:"audios,omitempty" xml:"audios>song_audios_dto,omitempty"`
	// 请求结果信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentSongsAudioGetrefrainAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Audios = m.Audios[:0]
	m.ResultCode = nil
}

var poolXiamiContentSongsAudioGetrefrainAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentSongsAudioGetrefrainAPIResponse)
	},
}

// GetXiamiContentSongsAudioGetrefrainAPIResponse 从 sync.Pool 获取 XiamiContentSongsAudioGetrefrainAPIResponse
func GetXiamiContentSongsAudioGetrefrainAPIResponse() *XiamiContentSongsAudioGetrefrainAPIResponse {
	return poolXiamiContentSongsAudioGetrefrainAPIResponse.Get().(*XiamiContentSongsAudioGetrefrainAPIResponse)
}

// ReleaseXiamiContentSongsAudioGetrefrainAPIResponse 将 XiamiContentSongsAudioGetrefrainAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentSongsAudioGetrefrainAPIResponse(v *XiamiContentSongsAudioGetrefrainAPIResponse) {
	v.Reset()
	poolXiamiContentSongsAudioGetrefrainAPIResponse.Put(v)
}
