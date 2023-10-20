package xiamicontent

import (
	"encoding/xml"

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
