package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
虾米-风格，流派歌曲 APIResponse
alibaba.xiami.api.tag.genre.song.get

虾米-风格，流派歌曲
*/
type AlibabaXiamiApiTagGenreSongGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaXiamiApiTagGenreSongGetResponse `json:"alibaba_xiami_api_tag_genre_song_get_response,omitempty"` 
    AlibabaXiamiApiTagGenreSongGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiTagGenreSongGetResponse struct {

    // 风格歌曲
    
    Data  *struct {
        TagGenreSongresult  *TagGenreSongresult `json:"tag_genre_songresult,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaXiamiApiTagGenreSongGetResponse struct {

    // 风格歌曲
    Data   *TagGenreSongresult `json:"data,omitempty"`

}
