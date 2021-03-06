package taotv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselPlaylistGetAPIResponse 根据频道ID获取频道下节目单以及当前播放 API返回值
// taobao.taotv.carousel.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
type TaobaoTaotvCarouselPlaylistGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvCarouselPlaylistGetAPIResponseModel
}

// TaobaoTaotvCarouselPlaylistGetAPIResponseModel is 根据频道ID获取频道下节目单以及当前播放 成功返回结果
type TaobaoTaotvCarouselPlaylistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_carousel_playlist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvCarouselPlaylistGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
