package xiamicontent

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌单详情接口 API返回值 
xiami.content.songs.collect.get

根据歌单id，获取歌单详情
*/
type XiamiContentSongsCollectGetAPIResponse struct {
    model.CommonResponse
    XiamiContentSongsCollectGetResponse
}

// 获取歌单详情接口 成功返回结果
type XiamiContentSongsCollectGetResponse struct {
    XMLName xml.Name `xml:"xiami_content_songs_collect_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的歌曲信息
    Songs   *Page `json:"songs,omitempty" xml:"songs,omitempty"`
    // 系统自动生成
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 歌单详情
    Collect   *CollectDTO `json:"collect,omitempty" xml:"collect,omitempty"`
}
