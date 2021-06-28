package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取播单列表 APIResponse
taobao.taotv.video.playlist.all

根据牌照和视频源等获取播单列表
*/
type TaobaoTaotvVideoPlaylistAllAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"taotv_video_playlist_all_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoTaotvVideoPlaylistAllResult `json:"result,omitempty" xml:"