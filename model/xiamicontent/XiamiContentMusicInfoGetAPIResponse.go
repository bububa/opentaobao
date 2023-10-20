package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMusicInfoGetAPIResponse 获取音乐信息 API返回值
// xiami.content.music.info.get
//
// (批量)获取歌曲信息
type XiamiContentMusicInfoGetAPIResponse struct {
	model.CommonResponse
	XiamiContentMusicInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentMusicInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentMusicInfoGetAPIResponseModel).Reset()
}

// XiamiContentMusicInfoGetAPIResponseModel is 获取音乐信息 成功返回结果
type XiamiContentMusicInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_music_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 音乐信息
	MusicList []MusicDto `json:"music_list,omitempty" xml:"music_list>music_dto,omitempty"`
	// 错误信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentMusicInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MusicList = m.MusicList[:0]
	m.ResultCode = nil
}

var poolXiamiContentMusicInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentMusicInfoGetAPIResponse)
	},
}

// GetXiamiContentMusicInfoGetAPIResponse 从 sync.Pool 获取 XiamiContentMusicInfoGetAPIResponse
func GetXiamiContentMusicInfoGetAPIResponse() *XiamiContentMusicInfoGetAPIResponse {
	return poolXiamiContentMusicInfoGetAPIResponse.Get().(*XiamiContentMusicInfoGetAPIResponse)
}

// ReleaseXiamiContentMusicInfoGetAPIResponse 将 XiamiContentMusicInfoGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentMusicInfoGetAPIResponse(v *XiamiContentMusicInfoGetAPIResponse) {
	v.Reset()
	poolXiamiContentMusicInfoGetAPIResponse.Put(v)
}
