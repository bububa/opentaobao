package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ott播单 API返回值 
taobao.taotv.video.playlist.ottnav.get

根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
*/
type TaobaoTaotvVideoPlaylistOttnavGetAPIResponse struct {
    model.CommonResponse
    TaobaoTaotvVideoPlaylistOttnavGetResponse
}

// ott播单 成功返回结果
type TaobaoTaotvVideoPlaylistOttnavGetResponse struct {
    XMLName xml.Name `xml:"taotv_video_playlist_ottnav_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoTaotvVideoPlaylistOttnavGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
