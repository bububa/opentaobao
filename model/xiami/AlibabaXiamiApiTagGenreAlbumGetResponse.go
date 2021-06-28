package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐－风格，流派专辑列表 APIResponse
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表
*/
type AlibabaXiamiApiTagGenreAlbumGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_tag_genre_album_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 风格，流派专辑列表
    
    Data   *TagAlbumResult `json:"data,omitempty" xml:"