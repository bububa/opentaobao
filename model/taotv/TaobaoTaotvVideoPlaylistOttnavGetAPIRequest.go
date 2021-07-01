package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvVideoPlaylistOttnavGetAPIRequest
ott播单 API请求
taobao.taotv.video.playlist.ottnav.get

根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息 */
type TaobaoTaotvVideoPlaylistOttnavGetAPIRequest struct {
	model.Params
	// 播单id
	_playListId int64
	// 播单列表
	_playListNav []string
	// 系统信息
	_systemInfo string
}

// New
