package taotv

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTaotvCarouselPlaylistGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvCarouselPlaylistGetAPIResponseModel).Reset()
}

// TaobaoTaotvCarouselPlaylistGetAPIResponseModel is 根据频道ID获取频道下节目单以及当前播放 成功返回结果
type TaobaoTaotvCarouselPlaylistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_carousel_playlist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvCarouselPlaylistGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvCarouselPlaylistGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvCarouselPlaylistGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselPlaylistGetAPIResponse)
	},
}

// GetTaobaoTaotvCarouselPlaylistGetAPIResponse 从 sync.Pool 获取 TaobaoTaotvCarouselPlaylistGetAPIResponse
func GetTaobaoTaotvCarouselPlaylistGetAPIResponse() *TaobaoTaotvCarouselPlaylistGetAPIResponse {
	return poolTaobaoTaotvCarouselPlaylistGetAPIResponse.Get().(*TaobaoTaotvCarouselPlaylistGetAPIResponse)
}

// ReleaseTaobaoTaotvCarouselPlaylistGetAPIResponse 将 TaobaoTaotvCarouselPlaylistGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvCarouselPlaylistGetAPIResponse(v *TaobaoTaotvCarouselPlaylistGetAPIResponse) {
	v.Reset()
	poolTaobaoTaotvCarouselPlaylistGetAPIResponse.Put(v)
}
