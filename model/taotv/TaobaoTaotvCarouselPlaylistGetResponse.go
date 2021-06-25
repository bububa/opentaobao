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
    Response *TaobaoTaotvCarouselPlaylistGetResponse `json:"taobao_taotv_carousel_playlist_get_response,omitempty"`
}

type TaobaoTaotvCarouselPlaylistGetResponse struct {

    // result
    Result   *TaobaoTaotvCarouselPlaylistGetResult `json:"result,omitempty"`

}
