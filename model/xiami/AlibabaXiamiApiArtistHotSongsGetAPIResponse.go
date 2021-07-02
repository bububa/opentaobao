package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiArtistHotSongsGetAPIResponse 热门歌曲 API返回值
// alibaba.xiami.api.artist.hotSongs.get
//
// 热门歌曲
type AlibabaXiamiApiArtistHotSongsGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiArtistHotSongsGetAPIResponseModel
}

// AlibabaXiamiApiArtistHotSongsGetAPIResponseModel is 热门歌曲 成功返回结果
type AlibabaXiamiApiArtistHotSongsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_artist_hotSongs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *HotSongsResult `json:"data,omitempty" xml:"data,omitempty"`
}
