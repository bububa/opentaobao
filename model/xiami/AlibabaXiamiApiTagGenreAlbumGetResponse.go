package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐－风格，流派专辑列表 APIResponse
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表
*/
type AlibabaXiamiApiTagGenreAlbumGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiTagGenreAlbumGetResponse `json:"alibaba_xiami_api_tag_genre_album_get_response,omitempty"`
}

type AlibabaXiamiApiTagGenreAlbumGetResponse struct {

    // 风格，流派专辑列表
    Data   *TagAlbumResult `json:"data,omitempty"`

}
