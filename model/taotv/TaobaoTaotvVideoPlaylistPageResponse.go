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
    Response *TaobaoTaotvVideoPlaylistPageResponse `json:"taobao_taotv_video_playlist_page_response,omitempty"`
}

type TaobaoTaotvVideoPlaylistPageResponse struct {

    // 接口返回model
    Result   *TaobaoTaotvVideoPlaylistPageResult `json:"result,omitempty"`

}
