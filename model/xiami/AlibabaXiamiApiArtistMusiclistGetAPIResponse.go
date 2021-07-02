package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiArtistMusiclistGetAPIResponse 热门艺人 API返回值
// alibaba.xiami.api.artist.musiclist.get
//
// 热门艺人
type AlibabaXiamiApiArtistMusiclistGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiArtistMusiclistGetAPIResponseModel
}

// AlibabaXiamiApiArtistMusiclistGetAPIResponseModel is 热门艺人 成功返回结果
type AlibabaXiamiApiArtistMusiclistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_artist_musiclist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// artist_music_list_result
	ArtistMusicListResult *AlibabaXiamiApiArtistMusiclistGetStruct `json:"artist_music_list_result,omitempty" xml:"artist_music_list_result,omitempty"`
}
