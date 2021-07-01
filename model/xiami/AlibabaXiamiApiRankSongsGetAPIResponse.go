package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiRankSongsGetAPIResponse
排行榜歌曲获取 API返回值
alibaba.xiami.api.rank.songs.get

获取歌曲排行榜 */
type AlibabaXiamiApiRankSongsGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiRankSongsGetAPIResponseModel
}

// AlibabaXiamiApiRankSongsGetAPIResponseModel is 排行榜歌曲获取 成功返回结果
type AlibabaXiamiApiRankSongsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_rank_songs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *RankSongsData `json:"data,omitempty" xml:"data,omitempty"`
}
