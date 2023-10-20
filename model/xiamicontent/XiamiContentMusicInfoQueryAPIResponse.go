package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMusicInfoQueryAPIResponse 搜索音乐 API返回值
// xiami.content.music.info.query
//
// (批量)获取歌曲信息
type XiamiContentMusicInfoQueryAPIResponse struct {
	model.CommonResponse
	XiamiContentMusicInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentMusicInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentMusicInfoQueryAPIResponseModel).Reset()
}

// XiamiContentMusicInfoQueryAPIResponseModel is 搜索音乐 成功返回结果
type XiamiContentMusicInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_music_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 音乐信息
	MusicDtoPage *Page `json:"music_dto_page,omitempty" xml:"music_dto_page,omitempty"`
	// 结果code
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentMusicInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MusicDtoPage = nil
	m.ResultCode = nil
}

var poolXiamiContentMusicInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentMusicInfoQueryAPIResponse)
	},
}

// GetXiamiContentMusicInfoQueryAPIResponse 从 sync.Pool 获取 XiamiContentMusicInfoQueryAPIResponse
func GetXiamiContentMusicInfoQueryAPIResponse() *XiamiContentMusicInfoQueryAPIResponse {
	return poolXiamiContentMusicInfoQueryAPIResponse.Get().(*XiamiContentMusicInfoQueryAPIResponse)
}

// ReleaseXiamiContentMusicInfoQueryAPIResponse 将 XiamiContentMusicInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentMusicInfoQueryAPIResponse(v *XiamiContentMusicInfoQueryAPIResponse) {
	v.Reset()
	poolXiamiContentMusicInfoQueryAPIResponse.Put(v)
}
