package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvVideoPlaylistGetAPIRequest
根据频道ID获取频道下节目单以及当前播放 API请求
taobao.taotv.video.playlist.get

根据频道ID获取频道下节目单以及当前播放 */
type TaobaoTaotvVideoPlaylistGetAPIRequest struct {
	model.Params
	// 播单id
	_playListId int64
	// 系统信息
	_systemInfo string
}

// New
