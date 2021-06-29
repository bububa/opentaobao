package xiamiopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲详情 APIResponse
xiami.api.song.detail.get

获取歌曲详情
*/
type XiamiApiSongDetailGetAPIResponse struct {
    model.CommonResponse
    XiamiApiSongDetailGetResponse
}

type XiamiApiSongDetailGetResponse struct {
    XMLName xml.Name `xml:"xiami_api_song_detail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 歌曲信息
    
    SongDetailList   []SongDetailDo `json:"song_detail_list,omitempty" xml:"song_detail_list>song_detail_do,omitempty"`
    
    
}
