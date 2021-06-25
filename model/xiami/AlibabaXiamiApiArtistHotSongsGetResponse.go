package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
热门歌曲 APIResponse
alibaba.xiami.api.artist.hotSongs.get

热门歌曲
*/
type AlibabaXiamiApiArtistHotSongsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiArtistHotSongsGetResponse `json:"alibaba_xiami_api_artist_hotSongs_get_response,omitempty"`
}

type AlibabaXiamiApiArtistHotSongsGetResponse struct {

    // 返回结果
    Data   *HotSongsResult `json:"data,omitempty"`

}
