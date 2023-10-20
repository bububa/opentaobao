package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsInfoGetAPIResponse 获取歌曲信息 API返回值
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
type XiamiContentSongsInfoGetAPIResponse struct {
	model.CommonResponse
	XiamiContentSongsInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentSongsInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentSongsInfoGetAPIResponseModel).Reset()
}

// XiamiContentSongsInfoGetAPIResponseModel is 获取歌曲信息 成功返回结果
type XiamiContentSongsInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲信息
	Songs []SongInfoDto `json:"songs,omitempty" xml:"songs>song_info_dto,omitempty"`
	// 系统自动生成
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentSongsInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Songs = m.Songs[:0]
	m.ResultCode = nil
}

var poolXiamiContentSongsInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentSongsInfoGetAPIResponse)
	},
}

// GetXiamiContentSongsInfoGetAPIResponse 从 sync.Pool 获取 XiamiContentSongsInfoGetAPIResponse
func GetXiamiContentSongsInfoGetAPIResponse() *XiamiContentSongsInfoGetAPIResponse {
	return poolXiamiContentSongsInfoGetAPIResponse.Get().(*XiamiContentSongsInfoGetAPIResponse)
}

// ReleaseXiamiContentSongsInfoGetAPIResponse 将 XiamiContentSongsInfoGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentSongsInfoGetAPIResponse(v *XiamiContentSongsInfoGetAPIResponse) {
	v.Reset()
	poolXiamiContentSongsInfoGetAPIResponse.Put(v)
}
