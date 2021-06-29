package xiamicontent

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲信息 APIResponse
xiami.content.songs.info.get

(批量)获取歌曲信息
*/
type XiamiContentSongsInfoGetAPIResponse struct {
    model.CommonResponse
    XiamiContentSongsInfoGetResponse
}

type XiamiContentSongsInfoGetResponse struct {
    XMLName xml.Name `xml:"xiami_content_songs_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 歌曲信息
    
    Songs   []SongInfoDto `json:"songs,omitempty" xml:"songs>song_info_dto,omitempty"`
    
    
}
