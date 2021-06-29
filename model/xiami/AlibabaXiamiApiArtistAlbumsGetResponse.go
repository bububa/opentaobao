package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
艺人专辑 APIResponse
alibaba.xiami.api.artist.albums.get

艺人专辑
*/
type AlibabaXiamiApiArtistAlbumsGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiArtistAlbumsGetResponse
}

type AlibabaXiamiApiArtistAlbumsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_artist_albums_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 艺人专辑结果
    
    Data   *ArtistAlbumsResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
