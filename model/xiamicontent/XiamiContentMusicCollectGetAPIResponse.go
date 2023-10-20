package xiamicontent

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMusicCollectGetAPIResponse 获取歌单歌曲 API返回值
// xiami.content.music.collect.get
//
// 获取歌单歌曲
type XiamiContentMusicCollectGetAPIResponse struct {
	model.CommonResponse
	XiamiContentMusicCollectGetAPIResponseModel
}

// XiamiContentMusicCollectGetAPIResponseModel is 获取歌单歌曲 成功返回结果
type XiamiContentMusicCollectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_music_collect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 音乐信息
	MusicDtoPage *Page `json:"music_dto_page,omitempty" xml:"music_dto_page,omitempty"`
	// 结果code
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
