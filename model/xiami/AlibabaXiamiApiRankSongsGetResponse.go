package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜歌曲获取 APIResponse
alibaba.xiami.api.rank.songs.get

获取歌曲排行榜
*/
type AlibabaXiamiApiRankSongsGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiRankSongsGetResponse
}

type AlibabaXiamiApiRankSongsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_rank_songs_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Data   *RankSongsData `json:"data,omitempty" xml:"data,omitempty"`

    
}
