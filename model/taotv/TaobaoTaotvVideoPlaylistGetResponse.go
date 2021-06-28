package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据频道ID获取频道下节目单以及当前播放 APIResponse
taobao.taotv.video.playlist.get

根据频道ID获取频道下节目单以及当前播放
*/
type TaobaoTaotvVideoPlaylistGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaotvVideoPlaylistGetResponse `json:"taotv_video_playlist_get_response,omitempty"` 
    TaobaoTaotvVideoPlaylistGetResponse
}

/* model for simplify = false
type TaobaoTaotvVideoPlaylistGetResponse struct {

    // result
    
    Result  *struct {
        TaobaoTaotvVideoPlaylistGetResult  *TaobaoTaotvVideoPlaylistGetResult `json:"taobao_taotv_video_playlist_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTaotvVideoPlaylistGetResponse struct {

    // result
    Result   *TaobaoTaotvVideoPlaylistGetResult `json:"result,omitempty"`

}
