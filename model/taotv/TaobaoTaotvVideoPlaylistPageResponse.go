package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取所有播单 API返回值 
taobao.taotv.video.playlist.page

获取所有播单信息（分页）
*/
type TaobaoTaotvVideoPlaylistPageAPIResponse struct {
    model.CommonResponse
    TaobaoTaotvVideoPlaylistPageResponse
}

// 分页获取所有播单 成功返回结果
type TaobaoTaotvVideoPlaylistPageResponse struct {
    XMLName xml.Name `xml:"taotv_video_playlist_page_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoTaotvVideoPlaylistPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
