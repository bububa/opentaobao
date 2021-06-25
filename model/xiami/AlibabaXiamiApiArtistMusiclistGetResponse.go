package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
热门艺人 APIResponse
alibaba.xiami.api.artist.musiclist.get

热门艺人
*/
type AlibabaXiamiApiArtistMusiclistGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiArtistMusiclistGetResponse `json:"alibaba_xiami_api_artist_musiclist_get_response,omitempty"`
}

type AlibabaXiamiApiArtistMusiclistGetResponse struct {

    // artist_music_list_result
    ArtistMusicListResult   *AlibabaXiamiApiArtistMusiclistGetStruct `json:"artist_music_list_result,omitempty"`

}
