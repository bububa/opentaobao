package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取播单列表 APIResponse
taobao.taotv.video.playlist.all

根据牌照和视频源等获取播单列表
*/
type TaobaoTaotvVideoPlaylistAllAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaotvVideoPlaylistAllResponse `json:"taotv_video_playlist_all_response,omitempty"` 
    TaobaoTaotvVideoPlaylistAllResponse
}

/* model for simplify = false
type TaobaoTaotvVideoPlaylistAllResponse struct {

    // result
    
    Result  *struct {
        TaobaoTaotvVideoPlaylistAllResult  *TaobaoTaotvVideoPlaylistAllResult `json:"taobao_taotv_video_playlist_all_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTaotvVideoPlaylistAllResponse struct {

    // result
    Result   *TaobaoTaotvVideoPlaylistAllResult `json:"result,omitempty"`

}
