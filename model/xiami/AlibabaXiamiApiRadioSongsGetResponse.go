package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取电台歌曲 APIResponse
alibaba.xiami.api.radio.songs.get

获取电台songs
*/
type AlibabaXiamiApiRadioSongsGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiRadioSongsGetResponse
}

type AlibabaXiamiApiRadioSongsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_radio_songs_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 歌曲列表
    
    Data   []StandardSong `json:"data,omitempty" xml:"data>standard_song,omitempty"`
    
    
}
