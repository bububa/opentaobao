package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
艺人专辑 APIResponse
alibaba.xiami.api.artist.albums.get

艺人专辑
*/
type AlibabaXiamiApiArtistAlbumsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiArtistAlbumsGetResponse `json:"alibaba_xiami_api_artist_albums_get_response,omitempty"`
}

type AlibabaXiamiApiArtistAlbumsGetResponse struct {

    // 艺人专辑结果
    Data   *ArtistAlbumsResult `json:"data,omitempty"`

}
