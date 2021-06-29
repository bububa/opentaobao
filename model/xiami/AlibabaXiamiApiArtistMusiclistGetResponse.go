package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
热门艺人 APIResponse
alibaba.xiami.api.artist.musiclist.get

热门艺人
*/
type AlibabaXiamiApiArtistMusiclistGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiArtistMusiclistGetResponse
}

type AlibabaXiamiApiArtistMusiclistGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_artist_musiclist_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // artist_music_list_result
    
    ArtistMusicListResult   *AlibabaXiamiApiArtistMusiclistGetStruct `json:"artist_music_list_result,omitempty" xml:"artist_music_list_result,omitempty"`

    
}
