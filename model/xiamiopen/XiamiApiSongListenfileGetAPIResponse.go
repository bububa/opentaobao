package xiamiopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiApiSongListenfileGetAPIResponse 获取歌曲试听文件 API返回值
// xiami.api.song.listenfile.get
//
// 获取歌曲试听文件
type XiamiApiSongListenfileGetAPIResponse struct {
	model.CommonResponse
	XiamiApiSongListenfileGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiApiSongListenfileGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiApiSongListenfileGetAPIResponseModel).Reset()
}

// XiamiApiSongListenfileGetAPIResponseModel is 获取歌曲试听文件 成功返回结果
type XiamiApiSongListenfileGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_api_song_listenfile_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲试听文件列表
	SongPlayInfoList []SongPlayInfoDo `json:"song_play_info_list,omitempty" xml:"song_play_info_list>song_play_info_do,omitempty"`
}

// Reset 清空结构体
func (m *XiamiApiSongListenfileGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SongPlayInfoList = m.SongPlayInfoList[:0]
}

var poolXiamiApiSongListenfileGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiApiSongListenfileGetAPIResponse)
	},
}

// GetXiamiApiSongListenfileGetAPIResponse 从 sync.Pool 获取 XiamiApiSongListenfileGetAPIResponse
func GetXiamiApiSongListenfileGetAPIResponse() *XiamiApiSongListenfileGetAPIResponse {
	return poolXiamiApiSongListenfileGetAPIResponse.Get().(*XiamiApiSongListenfileGetAPIResponse)
}

// ReleaseXiamiApiSongListenfileGetAPIResponse 将 XiamiApiSongListenfileGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiApiSongListenfileGetAPIResponse(v *XiamiApiSongListenfileGetAPIResponse) {
	v.Reset()
	poolXiamiApiSongListenfileGetAPIResponse.Put(v)
}
