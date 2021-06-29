package xiamicontent

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌单详情接口 APIResponse
xiami.content.songs.collect.get

根据歌单id，获取歌单详情
*/
type XiamiContentSongsCollectGetAPIResponse struct {
    model.CommonResponse
    XiamiContentSongsCollectGetResponse
}

type XiamiContentSongsCollectGetResponse struct {
    XMLName xml.Name `xml:"xiami_content_songs_collect_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的歌曲信息
    
    Songs   *Page `json:"songs,omitempty" xml:"songs,omitempty"`

    
    // 系统自动生成
    
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 歌单详情
    
    Collect   *CollectDto `json:"collect,omitempty" xml:"collect,omitempty"`

    
}
