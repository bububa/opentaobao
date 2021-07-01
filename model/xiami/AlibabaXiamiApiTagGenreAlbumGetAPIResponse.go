package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiTagGenreAlbumGetAPIResponse
虾米音乐－风格，流派专辑列表 API返回值
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表 */
type AlibabaXiamiApiTagGenreAlbumGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiTagGenreAlbumGetAPIResponseModel
}

// AlibabaXiamiApiTagGenreAlbumGetAPIResponseModel is 虾米音乐－风格，流派专辑列表 成功返回结果
type AlibabaXiamiApiTagGenreAlbumGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_tag_genre_album_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 风格，流派专辑列表
	Data *TagAlbumResult `json:"data,omitempty" xml:"data,omitempty"`
}
