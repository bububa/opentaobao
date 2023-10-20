package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsAudioGetAPIResponse 获取歌曲音频 API返回值
// xiami.content.songs.audio.get
//
// 获取歌曲音频
type XiamiContentSongsAudioGetAPIResponse struct {
	model.CommonResponse
	XiamiContentSongsAudioGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentSongsAudioGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentSongsAudioGetAPIResponseModel).Reset()
}

// XiamiContentSongsAudioGetAPIResponseModel is 获取歌曲音频 成功返回结果
type XiamiContentSongsAudioGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_audio_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 音频信息
	Audios []SongAudiosDto `json:"audios,omitempty" xml:"audios>song_audios_dto,omitempty"`
	// 请求结果信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentSongsAudioGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Audios = m.Audios[:0]
	m.ResultCode = nil
}

var poolXiamiContentSongsAudioGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentSongsAudioGetAPIResponse)
	},
}

// GetXiamiContentSongsAudioGetAPIResponse 从 sync.Pool 获取 XiamiContentSongsAudioGetAPIResponse
func GetXiamiContentSongsAudioGetAPIResponse() *XiamiContentSongsAudioGetAPIResponse {
	return poolXiamiContentSongsAudioGetAPIResponse.Get().(*XiamiContentSongsAudioGetAPIResponse)
}

// ReleaseXiamiContentSongsAudioGetAPIResponse 将 XiamiContentSongsAudioGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentSongsAudioGetAPIResponse(v *XiamiContentSongsAudioGetAPIResponse) {
	v.Reset()
	poolXiamiContentSongsAudioGetAPIResponse.Put(v)
}
