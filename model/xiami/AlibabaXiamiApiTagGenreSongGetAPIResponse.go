package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiTagGenreSongGetAPIResponse 虾米-风格，流派歌曲 API返回值
// alibaba.xiami.api.tag.genre.song.get
//
// 虾米-风格，流派歌曲
type AlibabaXiamiApiTagGenreSongGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiTagGenreSongGetAPIResponseModel
}

// AlibabaXiamiApiTagGenreSongGetAPIResponseModel is 虾米-风格，流派歌曲 成功返回结果
type AlibabaXiamiApiTagGenreSongGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_tag_genre_song_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 风格歌曲
	Data *TagGenreSongresult `json:"data,omitempty" xml:"data,omitempty"`
}
