package xiamiopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiApiSongDetailGetAPIResponse 获取歌曲详情 API返回值
// xiami.api.song.detail.get
//
// 获取歌曲详情
type XiamiApiSongDetailGetAPIResponse struct {
	model.CommonResponse
	XiamiApiSongDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiApiSongDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiApiSongDetailGetAPIResponseModel).Reset()
}

// XiamiApiSongDetailGetAPIResponseModel is 获取歌曲详情 成功返回结果
type XiamiApiSongDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_api_song_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲信息
	SongDetailList []SongDetailDo `json:"song_detail_list,omitempty" xml:"song_detail_list>song_detail_do,omitempty"`
}

// Reset 清空结构体
func (m *XiamiApiSongDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SongDetailList = m.SongDetailList[:0]
}

var poolXiamiApiSongDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiApiSongDetailGetAPIResponse)
	},
}

// GetXiamiApiSongDetailGetAPIResponse 从 sync.Pool 获取 XiamiApiSongDetailGetAPIResponse
func GetXiamiApiSongDetailGetAPIResponse() *XiamiApiSongDetailGetAPIResponse {
	return poolXiamiApiSongDetailGetAPIResponse.Get().(*XiamiApiSongDetailGetAPIResponse)
}

// ReleaseXiamiApiSongDetailGetAPIResponse 将 XiamiApiSongDetailGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiApiSongDetailGetAPIResponse(v *XiamiApiSongDetailGetAPIResponse) {
	v.Reset()
	poolXiamiApiSongDetailGetAPIResponse.Put(v)
}
