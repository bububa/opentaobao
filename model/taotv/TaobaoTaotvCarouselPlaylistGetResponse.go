package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据频道ID获取频道下节目单以及当前播放 APIResponse
taobao.taotv.carousel.playlist.get

根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
*/
type TaobaoTaotvCarouselPlaylistGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaotvCarouselPlaylistGetResponse `json:"taotv_carousel_playlist_get_response,omitempty"` 
    TaobaoTaotvCarouselPlaylistGetResponse
}

/* model for simplify = false
type TaobaoTaotvCarouselPlaylistGetResponse struct {

    // result
    
    Result  *struct {
        TaobaoTaotvCarouselPlaylistGetResult  *TaobaoTaotvCarouselPlaylistGetResult `json:"taobao_taotv_carousel_playlist_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTaotvCarouselPlaylistGetResponse struct {

    // result
    Result   *TaobaoTaotvCarouselPlaylistGetResult `json:"result,omitempty"`

}
