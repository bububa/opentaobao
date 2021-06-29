package xiamicontent

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索歌曲列表 APIResponse
xiami.content.songs.info.query

多维度查询歌曲列表
*/
type XiamiContentSongsInfoQueryAPIResponse struct {
    model.CommonResponse
    XiamiContentSongsInfoQueryResponse
}

type XiamiContentSongsInfoQueryResponse struct {
    XMLName xml.Name `xml:"xiami_content_songs_info_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的歌曲信息
    
    Songs   *Page `json:"songs,omitempty" xml:"songs,omitempty"`

    
    // 系统自动生成
    
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
