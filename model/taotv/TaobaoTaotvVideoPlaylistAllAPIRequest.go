package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvVideoPlaylistAllAPIRequest
获取播单列表 API请求
taobao.taotv.video.playlist.all

根据牌照和视频源等获取播单列表 */
type TaobaoTaotvVideoPlaylistAllAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// New
