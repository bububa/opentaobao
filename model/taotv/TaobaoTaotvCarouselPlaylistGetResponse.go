package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据频道ID获取频道下节目单以及当前播放 APIResponse
taobao.taotv.carousel.playlist.get

根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
*/
type TaobaoTaotvCarouselPlaylistGetAPIResponse struct {
    model.CommonResponse
    TaobaoTaotvCarouselPlaylistGetResponse
}

type TaobaoTaotvCarouselPlaylistGetResponse struct {
    XMLName xml.Name `xml:"taotv_carousel_playlist_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoTaotvCarouselPlaylistGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
