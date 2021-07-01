package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
艺人专辑 API返回值 
alibaba.xiami.api.artist.albums.get

艺人专辑
*/
type AlibabaXiamiApiArtistAlbumsGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiArtistAlbumsGetAPIResponseModel
}

// 艺人专辑 成功返回结果
type AlibabaXiamiApiArtistAlbumsGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_artist_albums_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 艺人专辑结果
    Data   *ArtistAlbumsResult `json:"data,omitempty" xml:"data,omitempty"`
}
