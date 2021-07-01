package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvCarouselPlaylistGetAPIRequest
根据频道ID获取频道下节目单以及当前播放 API请求
taobao.taotv.carousel.playlist.get

根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频 */
type TaobaoTaotvCarouselPlaylistGetAPIRequest struct {
	model.Params
	// 频道ID
	_channelId int64
	// 设备信息
	_systemInfo string
}

// New
