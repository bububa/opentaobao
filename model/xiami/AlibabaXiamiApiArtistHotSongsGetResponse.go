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
    // Response *AlibabaXiamiApiArtistHotSongsGetResponse `json:"alibaba_xiami_api_artist_hotSongs_get_response,omitempty"` 
    AlibabaXiamiApiArtistHotSongsGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiArtistHotSongsGetResponse struct {

    // 返回结果
    
    Data  *struct {
        HotSongsResult  *HotSongsResult `json:"hot_songs_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaXiamiApiArtistHotSongsGetResponse struct {

    // 返回结果
    Data   *HotSongsResult `json:"data,omitempty"`

}
