package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页获取所有播单 APIResponse
taobao.taotv.video.playlist.page

获取所有播单信息（分页）
*/
type TaobaoTaotvVideoPlaylistPageAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaotvVideoPlaylistPageResponse `json:"taotv_video_playlist_page_response,omitempty"` 
    TaobaoTaotvVideoPlaylistPageResponse
}

/* model for simplify = false
type TaobaoTaotvVideoPlaylistPageResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoTaotvVideoPlaylistPageResult  *TaobaoTaotvVideoPlaylistPageResult `json:"taobao_taotv_video_playlist_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTaotvVideoPlaylistPageResponse struct {

    // 接口返回model
    Result   *TaobaoTaotvVideoPlaylistPageResult `json:"result,omitempty"`

}
